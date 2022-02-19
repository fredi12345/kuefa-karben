package template

import (
	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type tmplBase struct {
	Authenticated bool
	Message       *Message
	PageLocation  string
}

func BaseTemplate(sess *sessions.Session, location string) *tmplBase {
	var t tmplBase
	t.initBase(sess, location)
	return &t
}

func (b *tmplBase) initBase(sess *sessions.Session, location string) {
	if auth, ok := sess.Values["authenticated"].(bool); ok {
		b.Authenticated = auth
	}

	// if there are multiple messages, only the last one is displayed
	// the other messages are discarded
	if flashes := sess.Flashes(); len(flashes) > 0 {
		if msg, ok := flashes[0].(*Message); ok {
			b.Message = msg
		}
	}

	b.PageLocation = location
}

type tmplEvent struct {
	Event *storage.Event
}

func (e *tmplEvent) initEvent(id string, service storage.Service) error {
	event, err := service.GetEvent(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get event "+id)
	}

	e.Event = event
	return nil
}

type tmplEventList struct {
	EventList []*storage.Event
}

func (el *tmplEventList) initEventList(page, cap int, service storage.Service) error {
	events, err := service.GetEventList(page, cap)
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
