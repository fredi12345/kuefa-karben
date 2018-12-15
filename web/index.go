package web

import (
	"net/http"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	templ, err := s.createIndexTmpl(sess)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	t := s.tmpl.Lookup("index.html")
	err = t.Execute(w, templ)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) createIndexTmpl(sess *sessions.Session) (tmplIndex, error) {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	events, err := s.db.GetEventList(1)
	if err != nil {
		panic(err)
	}

	length := len(events)
	if length > 2 {
		events = []*storage.Event{events[0], events[1]}
	}

	tmpl := tmplIndex{
		Authenticated: authenticated,
		EventList:     events,
	}

	if flashes := sess.Flashes(); len(flashes) > 0 {
		if msg, ok := flashes[0].(*message); ok {
			tmpl.Message = msg
		}
	}

	return tmpl, nil
}

type tmplIndex struct {
	Authenticated bool
	Message       *message
	EventList     []*storage.Event
}
