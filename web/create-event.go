package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

func (s *Server) CreateEventPage(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	templ := s.createCreateEventTmpl(sess)

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("create-event.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
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
