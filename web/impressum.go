package web

import (
	"net/http"

	"github.com/fredi12345/kuefa-karben/web/template"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) Impressum(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	tmpl := template.BaseTemplate(sess, "impressum")

	err := sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("impressum.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
