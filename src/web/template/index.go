package template

import (
	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type tmplIndex struct {
	tmplBase
	tmplEventList
	Comments     []*storage.Comment
	Participants []*storage.Participant
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

	newActivitiesLimit := 10
	tmpl.Participants, err = service.GetNewParticipants(newActivitiesLimit)
	if err != nil {
		return errors.WithMessage(err, "cannot get new participants for index")
	}
	tmpl.Comments, err = service.GetNewComments(newActivitiesLimit)
	if err != nil {
		return errors.WithMessage(err, "cannot get new comments for index")
	}

	return nil
}
