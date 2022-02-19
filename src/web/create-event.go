package web

import (
	"github.com/fredi12345/kuefa-karben/src/web/template"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) CreateEventPage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	tmpl := template.BaseTemplate(sess, "create-event")

	err := sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("create-event.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
