package web

import (
	"fmt"
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

//TODO: timeCreated und eventID speichern, siehe dazu mysql.go
func (s *Server) AddParticipant(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	fmt.Println("Request incoming")
	r.ParseForm()

	//TODO: Die Parameter kann man bestimmt einfacher zu int casten
	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}
	menu, err := strconv.Atoi(r.Form.Get("menu"))
	if err != nil {
		return err
	}
	name := r.Form.Get("name")

	part := storage.Participant{
		Name:    name,
		EventId: eventId,
		Menu:    menu,
		Created: time.Now()}

	err = s.db.CreateParticipant(part)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

	return nil
}
