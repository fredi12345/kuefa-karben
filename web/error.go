package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v", err)
	}

	templ := s.createNotFoundTmpl(sess)

	err = sess.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v", err)
	}

	t := s.tmpl.Lookup("not-found.html")
	err = t.Execute(w, templ)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v", err)
		w.WriteHeader(500)
	}
}

func (s *Server) createNotFoundTmpl(sess *sessions.Session) tmplNotFound {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	tmpl := tmplNotFound{
		Authenticated: authenticated,
		PageLocation:  "notFound",
	}

	return tmpl
}

type tmplNotFound struct {
	Authenticated bool
	PageLocation  string
	Message       *message
}

func (s *Server) HandleError(handler ErrorHandlerFunc) SessionHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, sess *sessions.Session) {
		err := handler(w, r, sess)

		if err != nil {
			if err == ErrWrongPassword || err == ErrNoAuthentication {
				redirectToIndex(sess, err, r, w)
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "error: %+v", err)
				redirectToIndex(sess, errors.New("Es ist ein unbekannter Fehler aufgetreten."), r, w)
			}
		}
	}
}

func redirectToIndex(sess *sessions.Session, err error, r *http.Request, w http.ResponseWriter) {
	sess.AddFlash(&message{Type: TypeError, Text: err.Error()})

	_ = sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
