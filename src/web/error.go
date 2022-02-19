package web

import (
	"fmt"
	template2 "github.com/fredi12345/kuefa-karben/src/web/template"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
	}

	tmpl := template2.BaseTemplate(sess, "notFound")

	err = sess.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
	}

	t := s.tmpl.Lookup("not-found.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		w.WriteHeader(500)
	}
}

func (s *Server) HandleError(handler ErrorHandlerFunc) SessionHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, sess *sessions.Session) {
		err := handler(w, r, sess)

		if err != nil {
			if err == ErrWrongPassword || err == ErrNoAuthentication {
				redirectToIndex(sess, err, r, w)
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
				redirectToIndex(sess, errors.New("Es ist ein unbekannter Fehler aufgetreten."), r, w)
			}
		}
	}
}

func redirectToIndex(sess *sessions.Session, err error, r *http.Request, w http.ResponseWriter) {
	sess.AddFlash(&template2.Message{Type: template2.TypeError, Text: err.Error()})

	_ = sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
