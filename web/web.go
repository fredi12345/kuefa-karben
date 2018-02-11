package web

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/storage"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type server struct {
	db storage.Service
}

func NewServer(db storage.Service) *server {
	return &server{db: db}
}

func (s *server) Index(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles(path.Join("resources", "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	id := 1

	ev, err := s.db.GetEvent(id)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, ev)
}

func (s *server) Participate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming")
	r.ParseForm()
	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		log.Fatal(err)
	}
	name := r.Form.Get("name")

	part := storage.Participant{
		Name:    name,
		EventId: eventId}

	err = s.db.CreateParticipant(part)
	if err != nil {
		log.Fatal(err)
	}
}
