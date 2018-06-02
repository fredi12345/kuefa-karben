package web

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/storage"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) DeleteParticipant(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	commentId, err := strconv.Atoi(r.Form.Get("participantId"))
	if err != nil {
		panic(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	err = s.db.DeleteParticipant(commentId)
	if err != nil {
		panic(err)
	}

	s.redirectToEventId(w, r, eventId)

}

//TODO: timeCreated und eventID speichern, siehe dazu mysql.go
func (s *Server) AddParticipant(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming")
	r.ParseForm()

	//TODO: Die Parameter kann man bestimmt einfacher zu int casten
	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}
	menu, err := strconv.Atoi(r.Form.Get("menu"))
	if err != nil {
		panic(err)
	}
	name := r.Form.Get("name")

	part := storage.Participant{
		Name:    name,
		EventId: eventId,
		Menu:    menu,
		Created: time.Now()}

	err = s.db.CreateParticipant(part)
	if err != nil {
		panic(err)
	}

	s.redirectToEventId(w, r, eventId)
}
