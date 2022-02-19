package template

import (
	"time"

	"github.com/fredi12345/kuefa-karben/src/storage"

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
	IsUpcoming           bool
	CommentsAllowed      bool
	Classic              int
	Vegetarian           int
	Vegan                int
}

func EventDetailTemplate(id string, sess *sessions.Session, service storage.Service) (*tmplEventDetail, error) {
	var t tmplEventDetail
	err := t.initTemplate(id, sess, service)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (tmpl *tmplEventDetail) initTemplate(id string, sess *sessions.Session, service storage.Service) error {
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
		return errors.WithMessage(err, "cannot get participants of event "+id)
	}
	tmpl.Participants = part
	classic, vegetarian, vegan := 0, 0, 0
	for i := 0; i < len(part); i++ {
		classic += part[i].ClassicCount
		vegetarian += part[i].VegetarianCount
		vegan += part[i].VeganCount
	}
	tmpl.Classic = classic
	tmpl.Vegetarian = vegetarian
	tmpl.Vegan = vegan

	imagesFileNames, err := service.GetImages(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get images of event "+id)
	}
	tmpl.ImageNames = imagesFileNames

	comments, err := service.GetComments(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get comments of event "+id)
	}
	tmpl.Comments = comments

	_, offset := time.Now().Zone()
	tmpl.ParticipationAllowed = time.Now().Before(tmpl.Event.ClosingDate.Add(time.Duration(-offset * 1000 * 1000 * 1000)))
	tmpl.IsUpcoming = time.Now().Before(tmpl.Event.EventDate.Add(time.Duration(-offset * 1000 * 1000 * 1000)))
	tmpl.CommentsAllowed = true

	return nil
}

func (tmpl *tmplEventDetail) SumParticipants() int {
	return tmpl.Classic + tmpl.Vegetarian + tmpl.Vegan
}
