package template

import (
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

type tmplIndex struct {
	tmplBase
	tmplEventList
}

func IndexTemplate(sess *sessions.Session, service storage.Service) (*tmplIndex, error) {
	var t tmplIndex
	err := t.initTemplate(sess, service)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (tmpl *tmplIndex) initTemplate(sess *sessions.Session, service storage.Service) error {
	tmpl.initBase(sess, "index")

	err := tmpl.initEventList(1, 2, service)
	if err != nil {
		return err
	}

	return nil
}
