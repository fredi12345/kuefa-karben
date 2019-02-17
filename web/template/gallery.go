package template

import (
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type tmplGallery struct {
	tmplBase
	PreviousPage int
	NextPage     int
	ImageNames   []*storage.Image
}

func GalleryTemplate(page int, maxPage int, imagesPerSite int, sess *sessions.Session, service storage.Service) (*tmplGallery, error) {
	var t tmplGallery
	err := t.initTemplate(page, maxPage, imagesPerSite, sess, service)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *tmplGallery) initTemplate(page int, maxPage int, imagesPerSite int, sess *sessions.Session, service storage.Service) error {
	t.initBase(sess, "gallery")

	if page <= 1 {
		t.PreviousPage = -1 // Im Template: Wenn <0 werden die Buttons ausgeblendet
	} else {
		t.PreviousPage = page - 1
	}

	if page < maxPage {
		t.NextPage = page + 1
	} else {
		t.NextPage = -1
	}

	imagesFileNames, err := service.GetAllImages(page, imagesPerSite)
	if err != nil {
		return errors.WithMessage(err, "cannot get images for gallery")
	}
	t.ImageNames = imagesFileNames

	return nil
}
