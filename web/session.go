package web

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) WithSession(handler SessionHandlerFunc) ErrorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		sess, err := s.cs.Get(r, cookieName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		return handler(w, r, sess)
	}
}

func (s *Server) HandleError(handler ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)

		if err != nil {
			panic(err)
		}
	}
}
