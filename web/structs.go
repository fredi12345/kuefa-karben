package web

import (
	"strconv"
	"time"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type tmplBase struct {
	Authenticated bool
	Message       *message
	PageLocation  string
}

func BaseTemplate(sess *sessions.Session, location string) *tmplBase {
	var t tmplBase
	t.initBase(sess, location)
	return &t
}

func (b *tmplBase) initBase(sess *sessions.Session, location string) {
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		b.Authenticated = auth
	}

	// if there are multiple messages, only the last one is displayed
	// the other messages are discarded
	if flashes := sess.Flashes(); len(flashes) > 0 {
		if msg, ok := flashes[0].(*message); ok {
			b.Message = msg
		}
	}

	b.PageLocation = location
}

type tmplEvent struct {
	Event *storage.Event
}

func (e *tmplEvent) initEvent(id int, service storage.Service) error {
	event, err := service.GetEvent(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get event "+strconv.Itoa(id))
	}

	e.Event = event
	return nil
}

type tmplEventList struct {
	EventList []*storage.Event
}

func (el *tmplEventList) initEventList(page, cap int, service storage.Service) error {
	events, err := service.GetEventList(page)
	if err != nil {
		return errors.WithMessage(err, "cannot get event list")
	}

	length := len(events)
	if length > cap {
		events = events[:cap] //*storage.Event{events[0], events[1]}
	}

	el.EventList = events
	return nil
}

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

type tmplAllEvents struct {
	tmplBase
	tmplEventList
	PreviousPage int
	NextPage     int
}

func AllEventsTemplate(page int, sess *sessions.Session, service storage.Service) (*tmplAllEvents, error) {
	var t tmplAllEvents
	t.initBase(sess, "eventList")

	err := t.initEventList(page, 9, service)
	if err != nil {
		return nil, err
	}

	if page <= 1 {
		t.PreviousPage = -1 // Im Template: Wenn <0 werden die Buttons ausgeblendet
	} else {
		t.PreviousPage = page - 1
	}

	eventCount, err := service.GetEventCount()
	if err != nil {
		return nil, errors.WithMessage(err, "cannot get event count")
	}

	if eventCount > page*9 {
		t.NextPage = page + 1
	} else {
		t.NextPage = -1
	}

	return &t, nil
}

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
	if err != nil { //TODO 404 statt unbekannter Fehler
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
