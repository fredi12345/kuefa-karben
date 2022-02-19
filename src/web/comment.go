package web

import (
	"fmt"
	"net/http"

	"github.com/fredi12345/kuefa-karben/src/storage"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	eventId := r.Form.Get("eventId")
	var c storage.Comment
	c.Name = r.Form.Get("name")
	c.Content = r.Form.Get("comment")
	c.EventID = eventId

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

	commentId := r.Form.Get("commentId")
	err = s.db.DeleteComment(commentId)
	if err != nil {
		return errors.WithMessage(err, "cannot delete comment "+commentId)
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#comments"), http.StatusSeeOther)
	return nil
}
