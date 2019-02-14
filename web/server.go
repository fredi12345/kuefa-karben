package web

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gorilla/securecookie"

	"github.com/fredi12345/kuefa-karben/random"
	"github.com/fredi12345/kuefa-karben/storage"
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

func NewServer(db storage.Service, imagePath string, thumbPath string, cookieKeyFile string) (*Server, error) {
	err := os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	t := template.Must(template.ParseGlob(path.Join("resources", "template", "**/*.tmpl")))
	t = template.Must(t.ParseGlob(path.Join("resources", "template", "*.html")))

	return &Server{
		db:        db,
		cs:        sessions.NewFilesystemStore("", getCookieKeys(cookieKeyFile)...),
		tmpl:      t,
		imgPath:   imagePath,
		thumbPath: thumbPath,
		rnd:       random.New(time.Now().UnixNano()),
	}, nil
}

func getCookieKeys(file string) [][]byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "generating new cookie keys...")
		return generateNewKeys(file)
	}

	return bytes.Split(b, []byte("\n"))
}

func generateNewKeys(file string) [][]byte {
	keys := [][]byte{securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32)}

	err := ioutil.WriteFile(file, bytes.Join(keys, []byte("\n")), 0666)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not write cookie key file: %v\n", err)
	}

	return keys
}

type SessionHandlerFunc func(w http.ResponseWriter, r *http.Request, session *sessions.Session)
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, session *sessions.Session) error
