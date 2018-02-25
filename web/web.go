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

	"io/ioutil"

	"net/url"

	"github.com/SchiffFlieger/go-random"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Server struct {
	db      storage.Service
	cs      sessions.Store
	tmpl    *template.Template
	imgPath string
	rnd     *random.Rnd
}

type templStruct struct {
	Event                *storage.Event
	Participants         []*storage.Participant
	ImageUrls            []string
	EventList            []*storage.Event
	Comments             []*storage.Comment
	Authenticated        bool
	ParticipationAllowed bool
	CommentsAllowed      bool
}

func (t *templStruct) HasImages() bool {
	return len(t.ImageUrls) > 0
}

func (t *templStruct) HasComments() bool {
	return len(t.Comments) > 0
}

const (
	cookieName = "session-cookie"
	cookieAuth = "authenticated"
)

func NewServer(db storage.Service, imagePath string) *Server {
	err := os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return &Server{
		db:      db,
		cs:      sessions.NewCookieStore(securecookie.GenerateRandomKey(64)),
		tmpl:    template.Must(template.ParseGlob(path.Join("resources", "template", "*.html"))),
		imgPath: imagePath,
		rnd:     random.New(time.Now().UnixNano()),
	}
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	id, err := s.getEventIdByUrl(r.URL)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	templ, err := s.createTemplateStruct(id, sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}

	t := s.tmpl.Lookup("index.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Impressum(w http.ResponseWriter, r *http.Request) {
	sess, err := s.cs.Get(r, cookieName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	templ, err := s.createTemplateStruct(1, sess)
	if err != nil {
		panic(err)
	}

	err = sess.Save(r, w)
	if err != nil {
		panic(err)
	}
	t := s.tmpl.Lookup("impressum.html")
	err = t.Execute(w, templ)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Comment(w http.ResponseWriter, r *http.Request) {
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

	http.Redirect(w, r, fmt.Sprintf("/?id=%d", eventId), http.StatusSeeOther)
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

	http.Redirect(w, r, fmt.Sprintf("/?id=%d", eventId), http.StatusSeeOther)
}

func (s *Server) createTemplateStruct(id int, sess *sessions.Session) (*templStruct, error) {
	var templ templStruct

	ev, err := s.db.GetEvent(id)
	if err != nil {
		return nil, err
	}
	templ.Event = ev

	part, err := s.db.GetParticipants(id)
	if err != nil {
		return nil, err
	}
	templ.Participants = part

	urls, err := s.db.GetImages(id)
	if err != nil {
		return nil, err
	}
	templ.ImageUrls = urls

	comments, err := s.db.GetComments(id)
	if err != nil {
		return nil, err
	}
	templ.Comments = comments

	events, err := s.db.GetEventList()
	if err != nil {
		return nil, err
	}
	templ.EventList = events

	templ.ParticipationAllowed = time.Now().Before(ev.EventDate)
	templ.CommentsAllowed = time.Now().After(ev.EventDate)

	if auth, ok := sess.Values[cookieAuth].(bool); ok && auth {
		templ.Authenticated = auth
	}

	return &templ, nil
}

func (s *Server) getEventIdByUrl(url *url.URL) (int, error) {
	keys, ok := url.Query()["id"]
	if !ok || len(keys) < 1 {
		return s.db.GetLatestEventId()
	}

	return strconv.Atoi(keys[0])
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

	http.Redirect(w, r, fmt.Sprintf("/?id=%d", eventId), http.StatusSeeOther)
}

func (s *Server) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		panic(err)
	}

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

	filename := s.writeFileToDisk(r)
	s.db.CreateImage("/public/images/"+filename, eventId)

	http.Redirect(w, r, fmt.Sprintf("/?id=%d", eventId), http.StatusSeeOther)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	user := r.Form.Get("user")
	pass := r.Form.Get("passwd")

	eventId, err := strconv.Atoi(r.Form.Get("eventId"))
	if err != nil {
		panic(err)
	}

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

	http.Redirect(w, r, fmt.Sprintf("/?id=%d", eventId), http.StatusSeeOther)
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		panic(err)
	}

	var event storage.Event
	event.Theme = r.Form.Get("theme")
	event.Starter = r.Form.Get("starter")
	event.MainDish = r.Form.Get("main-dish")
	event.Dessert = r.Form.Get("dessert")
	event.InfoText = r.Form.Get("info")

	d, err := time.Parse("2006-01-02", r.Form.Get("date"))
	if err != nil {
		panic(err)
	}
	event.EventDate = d

	filename := s.writeFileToDisk(r)
	event.ImageUrl = "/public/images/" + filename

	err = s.db.CreateEvent(event)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) writeFileToDisk(r *http.Request) string {
	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	infos, err := ioutil.ReadDir(s.imgPath)
	if err != nil {
		panic(err)
	}

	filename := strconv.Itoa(len(infos)) + handler.Filename

	f, err := os.Create(path.Join(s.imgPath, filename))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, file)
	if err != nil {
		panic(err)
	}

	return filename
}
