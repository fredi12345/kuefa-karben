package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func (s *Server) CreateEventPage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	templ := s.createCreateEventTmpl(sess)

	err := sess.Save(r, w)
	if err != nil {
		return err
	}

	t := s.tmpl.Lookup("create-event.html")
	err = t.Execute(w, templ)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) createCreateEventTmpl(sess *sessions.Session) tmplCreateEvent {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	return tmplCreateEvent{Authenticated: authenticated}
}

type tmplCreateEvent struct {
	Authenticated bool
}
