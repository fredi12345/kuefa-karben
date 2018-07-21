package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

func (s *Server) WithSession(handler SessionHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, err := s.cs.Get(r, cookieName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		handler(w, r, sess)
	}
}

func (s *Server) HandleError(handler ErrorHandlerFunc) SessionHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, sess *sessions.Session) {
		err := handler(w, r, sess)

		if err != nil {
			if err == ErrAuthenticationFailed {
				redirectToIndex(sess, err, r, w)
			} else {
				// TODO unknown error
				panic(err)
			}
		}
	}
}

func redirectToIndex(sess *sessions.Session, err error, r *http.Request, w http.ResponseWriter) {
	sess.AddFlash(&message{Type: TypeError, Text: err.Error()})

	err = sess.Save(r, w)
	if err != nil {
		panic(err) // TODO ist panic hier in Ordnung? was könnte man sonst machen? in welchem Fall schlägt save fehlt?
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
