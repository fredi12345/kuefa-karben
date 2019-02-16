package web

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"fmt"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/fredi12345/kuefa-karben/web/template"
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
	event.EventDate = d

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

	http.Redirect(w, r, fmt.Sprintf("/event/%d", id), http.StatusSeeOther)
	return nil
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	id, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return errors.WithStack(err)
	}

	images, err := s.db.GetImages(id)
	for _, image := range images {
		err = s.deleteImageById(image.Id)
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
		return errors.WithMessage(err, "cannot delete event "+strconv.Itoa(id))
	}
	sess.AddFlash(&template.Message{Type: template.TypeHint, Text: "Veranstaltung '" + event.Theme + "' erfolgreich gelöscht"})
	_ = sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func (s *Server) AllEvents(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		return errors.WithStack(err)
	}

	tmpl, err := template.AllEventsTemplate(page, sess, s.db)
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
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return errors.WithStack(err)
	}

	tmpl, err := template.EventDetailTemplate(id, sess, s.db)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			s.NotFound(w, r) // TODO ~KUF-61~ <- hat nichts damit zu tun? dieser fall funktioniert momentan nicht mehr
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

func (s *Server) getEventIdByUrl(url *url.URL) (int, error) {
	keys, ok := url.Query()["id"]
	if !ok || len(keys) < 1 {
		return s.db.GetLatestEventId()
	}

	return strconv.Atoi(keys[0])
}
