package web

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) DeleteParticipant(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	commentId, err := strconv.Atoi(r.Form.Get("participantId"))
	if err != nil {
		return err
	}

	err = s.db.DeleteParticipant(commentId)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}

func (s *Server) AddParticipant(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	r.ParseForm()

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}

	//Klassisch=0, Vegetarisch=1, Vegan=2
	menu, err := strconv.Atoi(r.Form.Get("menu"))
	if err != nil {
		return err
	}
	name := r.Form.Get("name")
	message := r.Form.Get("message")

	part := storage.Participant{
		Name:    name,
		EventId: eventId,
		Menu:    menu,
		Message: message,
		Created: time.Now()}

	err = s.db.CreateParticipant(part)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}
