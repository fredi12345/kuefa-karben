package web

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) WithSession(handler ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, err := s.cs.Get(r, cookieName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = handler(w, r, sess)
		if err != nil {
			panic(err)
		}
	}
}
