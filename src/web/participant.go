package web

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fredi12345/kuefa-karben/src/storage"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

func (s *Server) DeleteParticipant(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	err := r.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	participantId := r.Form.Get("participantId")

	err = s.db.DeleteParticipant(participantId)
	if err != nil {
		return errors.WithMessage(err, "cannot delete participant "+participantId)
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#participantList"), http.StatusSeeOther)
	return nil
}

func (s *Server) AddParticipant(w http.ResponseWriter, r *http.Request, sess *sessions.Session) error {
	r.ParseForm()

	eventId := r.Form.Get("eventId")
	classicCount, err := strconv.Atoi(r.Form.Get("classic_count"))
	vegetarianCount, err := strconv.Atoi(r.Form.Get("vegetarian_count"))
	veganCount, err := strconv.Atoi(r.Form.Get("vegan_count"))
	if err != nil {
		return errors.WithStack(err)
	}
	name := r.Form.Get("name")
	message := r.Form.Get("message")

	part := storage.Participant{
		Name:            name,
		EventID:         eventId,
		ClassicCount:    classicCount,
		VegetarianCount: vegetarianCount,
		VeganCount:      veganCount,
		Message:         message,
		Created:         time.Now()}

	err = s.db.CreateParticipant(part)
	if err != nil {
		return errors.WithMessage(err, "cannot create new participant")
	}

	http.Redirect(w, r, fmt.Sprint(r.Referer()+"#participation"), http.StatusSeeOther)
	return nil
}
