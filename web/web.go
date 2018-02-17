package web

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"

	"io"
	"os"

	"github.com/SchiffFlieger/go-random"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Server struct {
	db  storage.Service
	cs  sessions.Store
	rnd *random.Rnd
}

const (
	cookieName = "session-cookie"
	cookieAuth = "authenticated"
)

var templates *template.Template

func NewServer(db storage.Service) *Server {
	return &Server{
		db:  db,
		cs:  sessions.NewCookieStore(securecookie.GenerateRandomKey(64)),
		rnd: random.New(time.Now().UnixNano()),
	}
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	id := 1

	ev, err := s.db.GetEvent(id)
	if err != nil {
		panic(err)
	}

	part, err := s.db.GetParticipants(id)
	if err != nil {
		panic(err)
	}
	ev.Participants = part

	urls, err := s.db.GetImages(id)
	if err != nil {
		panic(err)
	}
	ev.ImageUrls = urls

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		ev.Authenticated = auth
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	templates, err = template.ParseGlob(path.Join("resources", "template", "*.html"))
	if err != nil {
		panic(err)
	}

	t := templates.Lookup("index.html")
	err = t.Execute(w, ev)
	if err != nil {
		panic(err)
	}
}

//TODO: timeCreated und eventID speichern, siehe dazu mysql.go
func (s *Server) Participate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming")
	r.ParseForm()

	//TODO: Die Parameter kann man bestimmt einfacher zu int casten
	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}
	menu, err := strconv.Atoi(r.Form.Get("menu"))
	if err != nil {
		panic(err)
	}
	name := r.Form.Get("name")

	part := storage.Participant{
		Name:    name,
		EventId: eventId,
		Menu:    menu,
		Created: time.Now()}

	err = s.db.CreateParticipant(part)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		panic(err)
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dest := path.Join("resources", "public", "images", handler.Filename)

	//TODO check that file names are unique
	f, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, file)
	if err != nil {
		panic(err)
	}

	s.db.CreateImage("/public/images/"+handler.Filename, 1)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	ok, err := s.db.CheckCredentials(user, pass)
	if err != nil {
		panic(err)
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		panic(err)
	}

	sess.Values[cookieAuth] = true

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
