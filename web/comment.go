package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/sessions"
)

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		return errors.WithStack(err)
	}

	var c storage.Comment
	c.Name = r.Form.Get("name")
	c.Content = r.Form.Get("comment")
	c.EventId = eventId

	err = s.db.CreateComment(c)
	if err != nil {
		return errors.WithMessage(err, "cannot create new comment")
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#comments"), http.StatusSeeOther)
	return nil
}

func (s *Server) DeleteComment(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	commentId, err := strconv.Atoi(r.Form.Get("commentId"))
	if err != nil {
		return errors.WithStack(err)
	}

	err = s.db.DeleteComment(commentId)
	if err != nil {
		return errors.WithMessage(err, "cannot delete comment "+strconv.Itoa(commentId))
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#comments"), http.StatusSeeOther)
	return nil
}
