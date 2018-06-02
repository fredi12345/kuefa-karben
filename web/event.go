package web

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (s *Server) Event(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
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

	s.redirectToEventId(w, r, redirectToLatest)
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

	s.redirectToEventId(w, r, redirectToLatest)
}

func (s *Server) getEventIdByUrl(url *url.URL) (int, error) {
	keys, ok := url.Query()["id"]
	if !ok || len(keys) < 1 {
		return s.db.GetLatestEventId()
	}

	return strconv.Atoi(keys[0])
}

func (s *Server) redirectToEventId(w http.ResponseWriter, r *http.Request, id int) {
	if id == redirectToLatest {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/?id=%d", id), http.StatusSeeOther)
	}
}

const (
	redirectToLatest = -1
)

type templStruct struct {
	Event                *storage.Event
	Participants         []*storage.Participant
	ImageUrls            []*storage.Image
	EventList            []*storage.Event
	Comments             []*storage.Comment
	Authenticated        bool
	ParticipationAllowed bool
	CommentsAllowed      bool
}

func (t *templStruct) HasImages() bool {
	return len(t.ImageUrls) > 0
}

func (t *templStruct) HasComments() bool {
	return len(t.Comments) > 0
}

func (s *Server) createTemplateStruct(id int, sess *sessions.Session) (*templStruct, error) {
	var templ templStruct

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
	templ.EventList = events

	templ.ParticipationAllowed = time.Now().Before(ev.EventDate)
	templ.CommentsAllowed = time.Now().After(ev.EventDate)

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		templ.Authenticated = auth
	}

	return &templ, nil
}
