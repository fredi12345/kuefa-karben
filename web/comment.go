package web

import (
	"net/http"
	"strconv"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return err
	}

	var c storage.Comment
	c.Name = r.Form.Get("name")
	c.Content = r.Form.Get("comment")
	c.EventId = eventId

	err = s.db.CreateComment(c)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

	return nil
}

func (s *Server) DeleteComment(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	commentId, err := strconv.Atoi(r.Form.Get("commentId"))
	if err != nil {
		return err
	}

	err = s.db.DeleteComment(commentId)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

	return nil
}
