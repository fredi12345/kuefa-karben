package web

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/fredi12345/kuefa-karben/src/storage"
	template2 "github.com/fredi12345/kuefa-karben/src/web/template"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return errors.WithStack(err)
	}

	var event storage.Event
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	if err != nil {
		return errors.WithStack(err)
	}
	c, err := time.Parse("2006-01-02T15:04", r.Form.Get("closingDate"))
	if err != nil {
		return errors.WithStack(err)
	}
	event.EventDate = d
	event.ClosingDate = c

	filename := getUniqueFileName()
	file, _, err := r.FormFile("image")
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()

	err = s.createAndSaveThumbAndFullImage(filename, file)
	if err != nil {
		return err
	}

	event.ImageName = filename

	id, err := s.db.CreateEvent(event)
	if err != nil {
		return errors.WithMessage(err, "cannot create new event")
	}

	http.Redirect(w, r, fmt.Sprintf("/event/%s", id), http.StatusSeeOther)
	return nil
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	id := r.Form.Get("eventId")
	images, err := s.db.GetImages(id)
	for _, image := range images {
		err = s.deleteImageById(image.ID)
	}

	event, err := s.db.GetEvent(id)
	if err != nil {
		return errors.WithMessage(err, "cannot get event to delete")
	}
	err = s.removeImageFileByFilename(event.ImageName)
	if err != nil {
		return err
	}

	err = s.db.DeleteEvent(id)
	if err != nil {
		return errors.WithMessage(err, "cannot delete event "+id)
	}
	sess.AddFlash(&template2.Message{Type: template2.TypeHint, Text: "Veranstaltung '" + event.Theme + "' erfolgreich gelöscht"})
	_ = sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (s *Server) AllEvents(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		return errors.WithStack(err)
	}

	const eventsPerSite = 9
	eventCount, err := s.db.GetEventCount()
	maxPage := int(math.Ceil(float64(eventCount) / eventsPerSite))

	if page > maxPage {
		http.Redirect(w, r, fmt.Sprintf("/event/all/%d", maxPage), http.StatusSeeOther)
		return nil
	}

	tmpl, err := template2.AllEventsTemplate(page, maxPage, eventsPerSite, sess, s.db)
	if err != nil {
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("event-all.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Server) EventDetail(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	id := mux.Vars(r)["id"]
	tmpl, err := template2.EventDetailTemplate(id, sess, s.db)
	if err != nil {
		if errors.Cause(err).Error() == "sql: no rows in result set" {
			s.NotFound(w, r)
			return nil
		}
		return err
	}

	err = sess.Save(r, w)
	if err != nil {
		return errors.WithStack(err)
	}

	t := s.tmpl.Lookup("event-detail.html")
	err = t.Execute(w, tmpl)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
