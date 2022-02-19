package template

import (
	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/gorilla/sessions"
)

type tmplEditEvent struct {
	tmplBase
	tmplEvent
}

func EditEventTemplate(id int, sess *sessions.Session, service storage.Service) (*tmplEditEvent, error) {
	var t tmplEditEvent
	err := t.initTemplate(id, sess, service)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *tmplEditEvent) initTemplate(id int, sess *sessions.Session, service storage.Service) error {
	t.initBase(sess, "edit-event")

	err := t.initEvent(id, service)
	return err
}
