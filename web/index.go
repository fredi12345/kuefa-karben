package web

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	id, err := s.getEventIdByUrl(r.URL)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	templ, err := s.createTemplateStruct(id, sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("index.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
}
