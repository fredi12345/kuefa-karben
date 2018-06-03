package web

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	templ, err := s.createIndexTmpl(sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("index.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
}

func (s *Server) createIndexTmpl(sess *sessions.Session) (tmplIndex, error) {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	events, err := s.db.GetEventList()
	if err != nil {
		panic(err)
	}

	length := len(events)
	if length > 2 {
		events = []*storage.Event{events[length-1], events[length-2]}
	}

	tmpl := tmplIndex{
		Authenticated: authenticated,
		EventList:     events,
	}

	return tmpl, nil
}

type tmplIndex struct {
	Authenticated bool
	EventList     []*storage.Event
}
