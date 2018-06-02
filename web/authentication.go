package web

import (
	"fmt"
	"os"
	"strconv"

	"net/http"
)

const (
	cookieName = "session-cookie"
	cookieAuth = "authenticated"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

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

	s.redirectToEventId(w, r, eventId)
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

	s.redirectToEventId(w, r, redirectToLatest)
}

func (s *Server) NeedsAuthentication(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, err := s.cs.Get(r, cookieName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
