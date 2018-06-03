package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"time"
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
	tmpl := tmplIndex{
		Authenticated: authenticated,
		EventList: []event{
			{
				Id:        1,
				EventDate: time.Now(),
				ImageUrl:  "/public/images/5progressLogo.Light.png",
				Theme:     "testtestasdf lecker",
			},
		},
	}

	return tmpl, nil
}

type tmplIndex struct {
	Authenticated bool
	EventList     []event
}

type event struct {
	Id        int
	Theme     string
	ImageUrl  string
	EventDate time.Time
}
