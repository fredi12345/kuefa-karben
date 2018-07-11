package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	templ, err := s.createNotFoundTmpl(sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("not-found.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
}

func (s *Server) createNotFoundTmpl(sess *sessions.Session) (tmplNotFound, error) {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	tmpl := tmplNotFound{
		Authenticated: authenticated,
	}

	return tmpl, nil
}

type tmplNotFound struct {
	Authenticated bool
}
