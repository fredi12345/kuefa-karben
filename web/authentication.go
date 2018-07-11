package web

import (
	"net/http"

	"fmt"

	"github.com/gorilla/sessions"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	ok, err := s.db.CheckCredentials(user, pass)
	if err != nil {
		panic(err)
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		panic(err)
	}

	sess.Values[cookieAuth] = true

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) NeedsAuthentication(handler ErrorHandlerFunc) ErrorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
		if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
			return handler(w, r, sess)
		}

		return fmt.Errorf("unauthorized")
	}
}
