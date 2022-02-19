package web

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) WithSession(handler SessionHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess, err := s.cs.Get(r, cookieName)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		}

		handler(w, r, sess)
	}
}
