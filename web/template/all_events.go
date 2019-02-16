package template

import (
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type tmplAllEvents struct {
	tmplBase
	tmplEventList
	PreviousPage int
	NextPage     int
}

func AllEventsTemplate(page int, sess *sessions.Session, service storage.Service) (*tmplAllEvents, error) {
	var t tmplAllEvents
	err := t.initTemplate(page, sess, service)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *tmplAllEvents) initTemplate(page int, sess *sessions.Session, service storage.Service) error {
	t.initBase(sess, "eventList")

	err := t.initEventList(page, 9, service)
	if err != nil {
		return err
	}

	if page <= 1 {
		t.PreviousPage = -1 // Im Template: Wenn <0 werden die Buttons ausgeblendet
	} else {
		t.PreviousPage = page - 1
	}

	eventCount, err := service.GetEventCount()
	if err != nil {
		return errors.WithMessage(err, "cannot get event count")
	}

	if eventCount > page*9 {
		t.NextPage = page + 1
	} else {
		t.NextPage = -1
	}

	return nil
}
