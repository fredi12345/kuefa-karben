package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fredi12345/kuefa-karben/src/storage"
	template2 "github.com/fredi12345/kuefa-karben/src/web/template"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) EditEventPage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	id := r.Form.Get("eventId")
	tmpl, err := template2.EditEventTemplate(id, sess, s.db)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("edit-event.html")
	err = t.Execute(w, tmpl)
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
	event.ID = r.Form.Get("event-id")
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	c, err := time.Parse("2006-01-02T15:04", r.Form.Get("closingDate"))
	if err != nil {
		return errors.WithStack(err)
	}
	event.EventDate = d
	event.ClosingDate = c

	file, header, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return errors.WithStack(err)
	}
	if file != nil {
		defer file.Close()
	}

	oldEvent, err := s.db.GetEvent(event.ID)
	if err != nil {
		return errors.WithMessage(err, "cannot get event "+event.ID)
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
		return errors.WithMessage(err, "cannot update event "+event.ID)
	}
	sess.AddFlash(&template2.Message{Type: template2.TypeHint, Text: "Veranstaltung erfolgreich bearbeitet"})
	_ = sess.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/event/%s", event.ID), http.StatusSeeOther)
	return nil
}
