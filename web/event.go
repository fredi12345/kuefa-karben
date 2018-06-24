package web

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func (s *Server) Event(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		panic(err)
	}

	var event storage.Event
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02", r.Form.Get("date"))
	if err != nil {
		panic(err)
	}
	event.EventDate = d

	filename := s.writeFileToDisk(r)
	event.ImageUrl = "/public/images/" + filename

	err = s.db.CreateEvent(event)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	err = s.db.DeleteEvent(eventId)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) AllEvents(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	tmpl, err := s.createTmplEventList(sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("event-all.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		panic(err)
	}
}

func (s *Server) EventDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	tmpl, err := s.createTmplEventDetail(id, sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("event-detail.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		panic(err)
	}
}

func (s *Server) getEventIdByUrl(url *url.URL) (int, error) {
	keys, ok := url.Query()["id"]
	if !ok || len(keys) < 1 {
		return s.db.GetLatestEventId()
	}

	return strconv.Atoi(keys[0])
}

const (
	redirectToLatest = -1
)

type tmplEventDetail struct {
	Event                *storage.Event
	Participants         []*storage.Participant
	ImageUrls            []*storage.Image
	EventList            []*storage.Event
	Comments             []*storage.Comment
	Authenticated        bool
	ParticipationAllowed bool
	CommentsAllowed      bool
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

	events, err := s.db.GetEventList()
	if err != nil {
		return nil, err
	}

	length := len(events)
	if length > 2 {
		events = []*storage.Event{events[length-1], events[length-2]}
	}

	templ.EventList = events

	templ.ParticipationAllowed = time.Now().Before(ev.EventDate)
	templ.CommentsAllowed = time.Now().After(ev.EventDate)

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		templ.Authenticated = auth
	}

	return &templ, nil
}

type tmplEventList struct {
	EventList     []*storage.Event
	Authenticated bool
}

func (s *Server) createTmplEventList(sess *sessions.Session) (*tmplEventList, error) {
	events, err := s.db.GetEventList()
	if err != nil {
		return nil, err
	}

	tmpl := tmplEventList{EventList: events}

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		tmpl.Authenticated = auth
	}

	return &tmpl, nil
}

func (s *Server) EditEvent(w http.ResponseWriter, r *http.Request) {

}
