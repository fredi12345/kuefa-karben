package web

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"bytes"
	"io"

	"fmt"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return err
	}

	var event storage.Event
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	if err != nil {
		return err
	}
	event.EventDate = d

	filename := getUniqueFileName()
	file, _, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()

	s.saveNewFile(file, filename)

	event.ImageUrl = "/public/images/" + filename

	id, err := s.db.CreateEvent(event)
	if err != nil {
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/event/%d", id), http.StatusSeeOther)

	return nil
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}

	err = s.db.DeleteEvent(eventId)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}

func (s *Server) AllEvents(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	tmpl, err := s.createTmplEventList(sess, r)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	t := s.tmpl.Lookup("event-all.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) EventDetail(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}

	tmpl, err := s.createTmplEventDetail(id, sess)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	t := s.tmpl.Lookup("event-detail.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) getEventIdByUrl(url *url.URL) (int, error) {
	keys, ok := url.Query()["id"]
	if !ok || len(keys) < 1 {
		return s.db.GetLatestEventId()
	}

	return strconv.Atoi(keys[0])
}

type tmplEventDetail struct {
	Event                *storage.Event
	Participants         []*storage.Participant
	ImageUrls            []*storage.Image
	EventList            []*storage.Event
	Comments             []*storage.Comment
	Authenticated        bool
	ParticipationAllowed bool
	CommentsAllowed      bool
	Message              *message
}

func (s *Server) createTmplEventDetail(id int, sess *sessions.Session) (*tmplEventDetail, error) {
	var templ tmplEventDetail

	ev, err := s.db.GetEvent(id)
	if err != nil {
		return nil, err
	}
	templ.Event = ev

	part, err := s.db.GetParticipants(id)
	if err != nil {
		return nil, err
	}
	templ.Participants = part

	urls, err := s.db.GetImages(id)
	if err != nil {
		return nil, err
	}
	templ.ImageUrls = urls

	comments, err := s.db.GetComments(id)
	if err != nil {
		return nil, err
	}
	templ.Comments = comments

	events, err := s.db.GetEventList(1)
	if err != nil {
		return nil, err
	}

	length := len(events)
	if length > 2 {
		events = []*storage.Event{events[0], events[1]}
	}

	templ.EventList = events

	templ.ParticipationAllowed = time.Now().Before(ev.EventDate)
	templ.CommentsAllowed = time.Now().After(ev.EventDate)

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		templ.Authenticated = auth
	}

	if flashes := sess.Flashes(); len(flashes) > 0 {
		if msg, ok := flashes[0].(*message); ok {
			templ.Message = msg
		}
	}

	return &templ, nil
}

type tmplEventList struct {
	EventList     []*storage.Event
	Authenticated bool
}

func (s *Server) createTmplEventList(sess *sessions.Session, r *http.Request) (*tmplEventList, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, err
	}

	events, err := s.db.GetEventList(id)
	if err != nil {
		return nil, err
	}

	tmpl := tmplEventList{EventList: events}

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		tmpl.Authenticated = auth
	}

	return &tmpl, nil
}

func (s *Server) EditEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return err
	}

	var event storage.Event
	id, err := strconv.Atoi(r.Form.Get("event-id"))
	if err != nil {
		return err
	}

	event.Id = id
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	if err != nil {
		return err
	}
	event.EventDate = d

	file, err := readFileFromRequest(r)
	if err != nil {
		return err
	}

	if file.Len() == 0 {
		err = s.db.UpdateEvent(event)
		if err != nil {
			return err
		}
	} else {
		s.updateEventWithNewImage(file, event)
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}

func (s *Server) updateEventWithNewImage(file *bytes.Buffer, event storage.Event) {
	filename := getUniqueFileName()
	err := s.saveNewFile(file, filename)
	if err != nil {
		panic(err)
	}
	err = s.db.UpdateEvent(event)
	if err != nil {
		panic(err)
	}
	err = s.db.UpdateEventImage(event.Id, "/public/images/"+filename)
	if err != nil {
		panic(err)
	}
}

func readFileFromRequest(r *http.Request) (*bytes.Buffer, error) {
	file, _, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
