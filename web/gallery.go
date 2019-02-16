package web

import (
	"github.com/fredi12345/kuefa-karben/web/template"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func (s *Server) Gallery(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	page, err := strconv.Atoi(mux.Vars(r)["page"])

	tmpl, err := template.GalleryTemplate(page, sess, s.db)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("gallery.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
