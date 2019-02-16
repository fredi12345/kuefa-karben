package web

import (
	"net/http"

	"github.com/fredi12345/kuefa-karben/web/template"
	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	templ, err := template.IndexTemplate(sess, s.db)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("index.html")
	err = t.Execute(w, templ)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
