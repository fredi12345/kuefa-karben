package web

import (
	"net/http"
	"time"

	"strconv"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) EditEventPage(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}

	templ := s.createEditEventTmpl(id, sess)

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	t := s.tmpl.Lookup("edit-event.html")
	err = t.Execute(w, templ)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) EditEvent(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		return err
	}

	var event storage.Event
	id, err := strconv.Atoi(r.Form.Get("event-id"))
	if err != nil {
		return err
	}

	event.Id = id
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02T15:04", r.Form.Get("date"))
	if err != nil {
		return err
	}
	event.EventDate = d

	file, header, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()

	if header.Size > 0 {
		filename, err := s.getEventImageName(id)
		if err != nil {
			return err
		}

		err = s.saveNewFullImageFile(file, filename)
		if err != nil {
			return err
		}
	}

	err = s.db.UpdateEvent(event)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	return nil
}

func (s *Server) getEventImageName(id int) (string, error) {
	ev, err := s.db.GetEvent(id)
	if err != nil {
		return "", err
	}

	return ev.ImageName, nil
}

func (s *Server) createEditEventTmpl(id int, sess *sessions.Session) tmplEditEvent {
	var authenticated bool
	if auth, ok := sess.Values[cookieAuth].(bool); ok {
		authenticated = auth
	}

	event, err := s.db.GetEvent(id)
	if err != nil {
		panic(err)
	}

	return tmplEditEvent{Authenticated: authenticated, PageLocation: "edit-event", Event: event}
}

type tmplEditEvent struct {
	Authenticated bool
	PageLocation  string
	Event         *storage.Event
}
