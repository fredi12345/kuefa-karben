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

				// Wert in session store schreiben
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				panic(err)
			}
		}
	}
}
