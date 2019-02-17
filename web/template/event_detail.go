package template

import (
	"strconv"
	"time"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type tmplEventDetail struct {
	tmplBase
	tmplEvent
	tmplEventList
	Participants         []*storage.Participant
	ImageNames           []*storage.Image
	Comments             []*storage.Comment
	ParticipationAllowed bool
	CommentsAllowed      bool
	Classic              int
	Vegetarian           int
	Vegan                int
}

func EventDetailTemplate(id int, sess *sessions.Session, service storage.Service) (*tmplEventDetail, error) {
	var t tmplEventDetail
	err := t.initTemplate(id, sess, service)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (tmpl *tmplEventDetail) initTemplate(id int, sess *sessions.Session, service storage.Service) error {
	tmpl.initBase(sess, "event")

	err := tmpl.initEvent(id, service)
	if err != nil {
		return err
	}

	err = tmpl.initEventList(1, 2, service)
	if err != nil {
		return err
	}

	part, err := service.GetParticipants(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get participants of event "+strconv.Itoa(id))
	}
	tmpl.Participants = part
	classic, vegetarian, vegan := 0, 0, 0
	for i := 0; i < len(part); i++ {
		switch part[i].Menu {
		case 0:
			classic++
			break
		case 1:
			vegetarian++
			break
		case 2:
			vegan++
			break
		}
	}
	tmpl.Classic = classic
	tmpl.Vegetarian = vegetarian
	tmpl.Vegan = vegan

	imagesFileNames, err := service.GetImages(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get images of event "+strconv.Itoa(id))
	}
	tmpl.ImageNames = imagesFileNames

	comments, err := service.GetComments(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get comments of event "+strconv.Itoa(id))
	}
	tmpl.Comments = comments

	tmpl.ParticipationAllowed = time.Now().Before(tmpl.Event.EventDate)
	tmpl.CommentsAllowed = true

	return nil
}
