package web

import (
	"net/http"

	"fmt"

	"github.com/gorilla/sessions"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	ok, err := s.db.CheckCredentials(user, pass)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("unauthorized")
	}

	sess.Values[cookieAuth] = true

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	sess.Values[cookieAuth] = false

	err := sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (s *Server) NeedsAuthentication(handler SessionHandlerFunc) SessionHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
		if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
			return handler(w, r, sess)
		}

		return fmt.Errorf("unauthorized")
	}
}
