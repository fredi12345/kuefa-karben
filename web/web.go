package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"io"
	"os"

	"github.com/fredi12345/kuefa-karben/storage"
)

type Server struct {
	db storage.Service
}

func NewServer(db storage.Service) *Server {
	return &Server{db: db}
}

func (s *Server) Index(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles(path.Join("resources", "template", "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	id := 1

	ev, err := s.db.GetEvent(id)
	if err != nil {
		log.Fatal(err)
	}

	part, err := s.db.GetParticipants(id)
	if err != nil {
		log.Fatal(err)
	}

	ev.Participants = part

	t.Execute(w, ev)
}

//TODO: timeCreated und eventID speichern, siehe dazu mysql.go
func (s *Server) Participate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming")
	r.ParseForm()

	//TODO: Die Parameter kann man bestimmt einfacher zu int casten
	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		log.Fatal(err)
	}
	menu, err := strconv.Atoi(r.Form.Get("menu"))
	if err != nil {
		log.Fatal(err)
	}
	name := r.Form.Get("name")

	part := storage.Participant{
		Name:    name,
		EventId: eventId,
		Menu:    menu,
		Created: time.Now()}

	err = s.db.CreateParticipant(part)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		log.Fatal(err)
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dest := path.Join("resources", "public", "images", handler.Filename)

	//TODO check that file names are unique
	f, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
	}

	s.db.CreateImage("/public/images/"+handler.Filename, 1)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
