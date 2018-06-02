package web

import (
	"github.com/fredi12345/kuefa-karben/storage"
	"net/http"
	"strconv"
)

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	var c storage.Comment
	c.Name = r.Form.Get("name")
	c.Content = r.Form.Get("comment")
	c.EventId = eventId

	err = s.db.CreateComment(c)
	if err != nil {
		panic(err)
	}

	s.redirectToEventId(w, r, eventId)
}

func (s *Server) DeleteComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	commentId, err := strconv.Atoi(r.Form.Get("commentId"))
	if err != nil {
		panic(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	err = s.db.DeleteComment(commentId)
	if err != nil {
		panic(err)
	}

	s.redirectToEventId(w, r, eventId)
}
