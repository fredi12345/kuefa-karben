package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func (s *Server) Impressum(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	//TODO: hier wird Event Detail template genutzt
	templ, err := s.createTmplEventDetail(1, sess)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return err
	}
	t := s.tmpl.Lookup("impressum.html")
	err = t.Execute(w, templ)
	if err != nil {
		return err
	}

	return nil
}
