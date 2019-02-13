package web

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"fmt"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return errors.WithStack(err)
	}

	var event storage.Event
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	if err != nil {
		return errors.WithStack(err)
	}
	event.EventDate = d

	filename := getUniqueFileName()
	file, _, err := r.FormFile("image")
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()

	err = s.createAndSaveThumbAndFullImage(filename, file)
	if err != nil {
		return err
	}

	event.ImageName = filename

	id, err := s.db.CreateEvent(event)
	if err != nil {
		return errors.Wrap(err, "cannot create new event")
	}

	http.Redirect(w, r, fmt.Sprintf("/event/%d", id), http.StatusSeeOther)
	return nil
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	id, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return errors.WithStack(err)
	}

	images, err := s.db.GetImages(id)
	for _, image := range images {
		err = s.deleteImageById(image.Id)
	}

	event, err := s.db.GetEvent(id)
	err = s.removeImageFileByFilename(event.ImageName)
	if err != nil {
		return err
	}

	err = s.db.DeleteEvent(id)
	if err != nil {
		return errors.Wrap(err, "cannot delete event "+strconv.Itoa(id))
	}
	sess.AddFlash(&message{Type: TypeHint, Text: "Veranstaltung '" + event.Theme + "' erfolgreich gelöscht"})
	_ = sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (s *Server) AllEvents(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	tmpl, err := s.createTmplEventList(sess, r)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("event-all.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Server) EventDetail(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return errors.WithStack(err)
	}

	tmpl, err := s.createTmplEventDetail(id, sess)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			s.NotFound(w, r) // TODO KUF-61
			return nil
		}
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("event-detail.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
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
	Authenticated        bool
	PageLocation         string
	Message              *message
	Event                *storage.Event
	Participants         []*storage.Participant
	ImageNames           []*storage.Image
	EventList            []*storage.Event
	Comments             []*storage.Comment
	ParticipationAllowed bool
	CommentsAllowed      bool
	Classic              int
	Vegetarian           int
	Vegan                int
}

func (s *Server) createTmplEventDetail(id int, sess *sessions.Session) (*tmplEventDetail, error) {
	var templ tmplEventDetail

	ev, err := s.db.GetEvent(id)
	if err != nil { //TODO 404 statt unbekannter Fehler
		return nil, errors.Wrap(err, "cannot get event "+strconv.Itoa(id))
	}
	templ.Event = ev

	part, err := s.db.GetParticipants(id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get participants of event "+strconv.Itoa(id))
	}
	templ.Participants = part
	classic, vegetarian, vegan := 0, 0, 0
	for i := 0; i < len(part); i++ {
		switch part[i].Menu {
		case 0:
			classic++
			break
		case 1:
			vegetarian++
			break
		case 2:
			vegan++
			break
		}
	}
	templ.Classic = classic
	templ.Vegetarian = vegetarian
	templ.Vegan = vegan

	imagesFileNames, err := s.db.GetImages(id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get images of event "+strconv.Itoa(id))
	}
	templ.ImageNames = imagesFileNames

	comments, err := s.db.GetComments(id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get comments of event "+strconv.Itoa(id))
	}
	templ.Comments = comments

	events, err := s.db.GetEventList(1)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get event list")
	}

	length := len(events)
	if length > 2 {
		events = []*storage.Event{events[0], events[1]}
	}

	templ.EventList = events

	templ.ParticipationAllowed = time.Now().Before(ev.EventDate)
	templ.CommentsAllowed = time.Now().After(ev.EventDate)
	templ.PageLocation = "event"

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
	Authenticated bool
	PageLocation  string
	Message       *message
	EventList     []*storage.Event
	PreviousPage  int
	NextPage      int
}

func (s *Server) createTmplEventList(sess *sessions.Session, r *http.Request) (*tmplEventList, error) {
	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	events, err := s.db.GetEventList(page)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get event list")
	}

	tmpl := tmplEventList{EventList: events, PageLocation: "eventList"}

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		tmpl.Authenticated = auth
	}

	if page <= 1 {
		tmpl.PreviousPage = -1 // Im Template: Wenn <0 werden die Buttons ausgeblendet
	} else {
		tmpl.PreviousPage = page - 1
	}

	eventCount, err := s.db.GetEventCount()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get event count")
	}

	if eventCount > page*9 {
		tmpl.NextPage = page + 1
	} else {
		tmpl.NextPage = -1
	}

	return &tmpl, nil
}
