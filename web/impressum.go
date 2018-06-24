package web

import (
	"fmt"
	"net/http"
	"os"
)

func (s *Server) Impressum(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	//TODO: hier wird Event Detail template genutzt
	templ, err := s.createTmplEventDetail(1, sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}
	t := s.tmpl.Lookup("impressum.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
}
