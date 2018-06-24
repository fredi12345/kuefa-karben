package web

import (
	"fmt"
	"net/http"
	"os"

	"strconv"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) EditEventPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	id, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	templ := s.createEditEventTmpl(id, sess)

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("edit-event.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
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

	return tmplEditEvent{Authenticated: authenticated, Event: event}
}

type tmplEditEvent struct {
	Authenticated bool
	Event         *storage.Event
}
