package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"strconv"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) EditEventPage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	id, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return errors.WithStack(err)
	}

	templ, err := s.createEditEventTmpl(id, sess)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("edit-event.html")
	err = t.Execute(w, templ)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Server) EditEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return errors.WithStack(err)
	}

	var event storage.Event
	id, err := strconv.Atoi(r.Form.Get("event-id"))
	if err != nil {
		return errors.WithStack(err)
	}

	event.Id = id
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	if err != nil {
		return errors.WithStack(err)
	}
	event.EventDate = d

	file, header, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return errors.WithStack(err)
	}
	if file != nil {
		defer file.Close()
	}

	oldEvent, err := s.db.GetEvent(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get event"+strconv.Itoa(id))
	}

	event.ImageName = oldEvent.ImageName

	if header != nil && header.Size > 0 {
		err = s.removeImageFileByFilename(oldEvent.ImageName)
		if err != nil {
			return err
		}
		filename := getUniqueFileName()
		err = s.createAndSaveThumbAndFullImage(filename, file)
		if err != nil {
			return err
		}

		event.ImageName = filename
	}

	err = s.db.UpdateEvent(event)
	if err != nil {
		return errors.WithMessage(err, "cannot update event "+strconv.Itoa(id))
	}
	sess.AddFlash(&message{Type: TypeHint, Text: "Veranstaltung erfolgreich bearbeitet"})
	_ = sess.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/event/%d", id), http.StatusSeeOther)
	return nil
}

func (s *Server) getEventImageName(id int) (string, error) {
	ev, err := s.db.GetEvent(id)
	if err != nil {
		return "", errors.WithMessage(err, "cannot get event "+strconv.Itoa(id))
	}

	return ev.ImageName, nil
}

func (s *Server) createEditEventTmpl(id int, sess *sessions.Session) (tmplEditEvent, error) {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	event, err := s.db.GetEvent(id)
	if err != nil {
		return tmplEditEvent{}, errors.WithMessage(err, "cannot get event "+strconv.Itoa(id))
	}

	return tmplEditEvent{Authenticated: authenticated, PageLocation: "edit-event", Event: event}, nil
}

type tmplEditEvent struct {
	Authenticated bool
	PageLocation  string
	Message       *message
	Event         *storage.Event
}
