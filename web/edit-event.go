package web

import (
	"net/http"

	"strconv"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) EditEventPage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}

	templ := s.createEditEventTmpl(id, sess)

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	t := s.tmpl.Lookup("edit-event.html")
	err = t.Execute(w, templ)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) createEditEventTmpl(id int, sess *sessions.Session) tmplEditEvent {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	event, err := s.db.GetEvent(id)
	if err != nil {
		panic(err)
	}

	return tmplEditEvent{Authenticated: authenticated, PageLocation: "edit-event", Event: event}
}

type tmplEditEvent struct {
	Authenticated bool
	PageLocation  string
	Event         *storage.Event
}
