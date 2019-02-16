package web

import (
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func (s *Server) Gallery(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	page, err := strconv.Atoi(mux.Vars(r)["page"])
	templ, err := s.createTmplGallery(sess, page)

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("gallery.html")
	err = t.Execute(w, templ)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type tmplGallery struct {
	Authenticated bool
	PageLocation  string
	Message       *message
	PreviousPage  int
	NextPage      int
	ImageNames    []*storage.Image
}

func (s *Server) createTmplGallery(sess *sessions.Session, page int) (tmplGallery, error) {
	templ := tmplGallery{PageLocation: "gallery"}

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		templ.Authenticated = auth
	}
	if page <= 1 {
		templ.PreviousPage = -1 // Im Template: Wenn <0 werden die Buttons ausgeblendet
	} else {
		templ.PreviousPage = page - 1
	}
	imageCount, err := s.db.GetImageCount()
	if err != nil {
		return tmplGallery{}, errors.WithMessage(err, "cannot get event count")
	}

	if imageCount > page*16 {
		templ.NextPage = page + 1
	} else {
		templ.NextPage = -1
	}

	imagesFileNames, err := s.db.GetAllImages(page)
	if err != nil {
		return tmplGallery{}, errors.WithMessage(err, "cannot get images for gallery")
	}
	templ.ImageNames = imagesFileNames

	return templ, err
}
