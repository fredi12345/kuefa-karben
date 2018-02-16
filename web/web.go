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

	"homespace/random"

	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Server struct {
	db storage.Service
	cs sessions.Store
}

const cookieName = "session-cookie"

func NewServer(db storage.Service) *Server {
	return &Server{db: db, cs: sessions.NewCookieStore(securecookie.GenerateRandomKey(64))}
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

	urls, err := s.db.GetImages(id)
	if err != nil {
		log.Fatal(err)
	}
	ev.ImageUrls = urls

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

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	ok, err := s.db.CheckCredentials(user, pass)
	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessName := random.Rnd.RandomString(32)
	sess, err := s.cs.New(r, sessName)
	if err != nil {
		log.Fatal(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}

	secureCookie := securecookie.New(securecookie.GenerateRandomKey(32), nil)
	encoded, err := secureCookie.Encode(cookieName, sessName)
	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    encoded,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
