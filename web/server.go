package web

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/fredi12345/kuefa-karben/random"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

const (
	cookieName = "session-cookie"
	cookieAuth = "authenticated"
)

type Server struct {
	db        storage.Service
	cs        sessions.Store
	tmpl      *template.Template
	imgPath   string
	thumbPath string
	rnd       *random.Rnd
}

func NewServer(db storage.Service, imagePath string, thumbPath string) (*Server, error) {
	err := os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	t := template.Must(template.ParseGlob(path.Join("resources", "template", "**/*.tmpl")))
	t = template.Must(t.ParseGlob(path.Join("resources", "template", "*.html")))

	return &Server{
		db:        db,
		cs:        sessions.NewCookieStore(securecookie.GenerateRandomKey(64)),
		tmpl:      t,
		imgPath:   imagePath,
		thumbPath: thumbPath,
		rnd:       random.New(time.Now().UnixNano()),
	}, nil
}

type SessionHandlerFunc func(w http.ResponseWriter, r *http.Request, session *sessions.Session)
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, session *sessions.Session) error
