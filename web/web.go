package web

import (
	"net/http"
	"path"
	"log"
	"github.com/fredi12345/kuefa-karben/storage"
	"html/template"
)

type server struct {
	db storage.Service
}

func NewServer(db storage.Service) *server  {
	return &server{db: db}
}

func (s *server) Index(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles(path.Join("resources", "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	ev, err := s.db.GetEvent(1)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, ev)
}
