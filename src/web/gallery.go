package web

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/src/web/template"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func (s *Server) Gallery(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	page, err := strconv.Atoi(mux.Vars(r)["page"])

	const imagesPerSite = 16
	imageCount, err := s.db.GetImageCount()
	maxPage := int(math.Ceil(float64(imageCount) / imagesPerSite))

	if page > maxPage {
		http.Redirect(w, r, fmt.Sprintf("/gallery/%d", maxPage), http.StatusSeeOther)
		return nil
	}

	tmpl, err := template.GalleryTemplate(page, maxPage, imagesPerSite, sess, s.db)
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
