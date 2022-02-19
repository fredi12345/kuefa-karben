package web

import (
	"net/http"

	"github.com/pkg/errors"

	"fmt"

	"github.com/gorilla/sessions"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	ok, err := s.db.CheckCredentials(user, pass)
	if err != nil || !ok {
		return ErrWrongPassword
	}

	sess.Values[cookieAuth] = true

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	return nil
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	sess.Values[cookieAuth] = false

	err := sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	return nil
}

func (s *Server) NeedsAuthentication(handler ErrorHandlerFunc) ErrorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
		if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
			return handler(w, r, sess)
		}

		return ErrNoAuthentication
	}
}

var (
	ErrWrongPassword    = fmt.Errorf("Nutzername oder Passwort falsch!")
	ErrNoAuthentication = fmt.Errorf("Sie haben nicht die erforderlichen Rechte um auf die Seite zuzugreifen.")
)
