package web

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) Impressum(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	templ := s.createTmplImpressum(sess)

	err := sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("impressum.html")
	err = t.Execute(w, templ)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type tmplImpressum struct {
	PageLocation  string
	Authenticated bool
}

func (s *Server) createTmplImpressum(sess *sessions.Session) tmplImpressum {
	var templ tmplImpressum
	templ.PageLocation = "impressum"

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		templ.Authenticated = auth
	}

	return templ
}
