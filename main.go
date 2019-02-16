package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/fredi12345/kuefa-karben/config"
	"github.com/fredi12345/kuefa-karben/storage/mydb"
	"github.com/fredi12345/kuefa-karben/web"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Read("config.xml")
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	db, err := mydb.New(cfg)
	if err != nil {
		log.Fatalf("could not create database: %v", err)
	}

	//TODO Ordner erstellen wenn nicht vorhanden
	imgPath := path.Join("resources", "public", "images")
	thumbPath := path.Join("resources", "public", "thumbs")

	server, err := web.NewServer(db, imgPath, thumbPath, "cookies.key")
	if err != nil {
		log.Fatal("could not create server: %v", err)
	}

	handler := createHandler(server)
	fmt.Println("http://localhost:63726")

	if err := http.ListenAndServe(":63726", handler); err != nil {
		log.Fatal(err)
	}
}

func createHandler(server *web.Server) http.Handler {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("resources/public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	// redirect from incomplete urls
	r.Handle("/event", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all/", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all/0", http.RedirectHandler("/event/all/1", http.StatusSeeOther))

	// http get methods
	r.HandleFunc("/", server.WithSession(server.HandleError(server.Index))).Methods(http.MethodGet)
	r.HandleFunc("/event/create", server.WithSession(server.HandleError(server.NeedsAuthentication(server.CreateEventPage)))).Methods(http.MethodGet)
	r.HandleFunc("/event/edit", server.WithSession(server.HandleError(server.NeedsAuthentication(server.EditEventPage)))).Methods(http.MethodGet)
	r.HandleFunc("/event/{id:[0-9]+}", server.WithSession(server.HandleError(server.EventDetail))).Methods(http.MethodGet)
	r.HandleFunc("/event/all/{page:[0-9]+}", server.WithSession(server.HandleError(server.AllEvents))).Methods(http.MethodGet)
	r.HandleFunc("/impressum", server.WithSession(server.HandleError(server.Impressum))).Methods(http.MethodGet)
	r.HandleFunc("/gallery/{page:[0-9]*}", server.WithSession(server.HandleError(server.Gallery))).Methods(http.MethodGet)

	// http post methods
	r.HandleFunc("/event/add", server.WithSession(server.HandleError(server.NeedsAuthentication(server.AddEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/event/edit", server.WithSession(server.HandleError(server.NeedsAuthentication(server.EditEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/event/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/participant/add", server.WithSession(server.HandleError(server.AddParticipant))).Methods(http.MethodPost)
	r.HandleFunc("/participant/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteParticipant)))).Methods(http.MethodPost)
	r.HandleFunc("/comment/add", server.WithSession(server.HandleError(server.AddComment))).Methods(http.MethodPost)
	r.HandleFunc("/comment/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteComment)))).Methods(http.MethodPost)
	r.HandleFunc("/image/add", server.WithSession(server.HandleError(server.NeedsAuthentication(server.AddImage)))).Methods(http.MethodPost)
	r.HandleFunc("/image/delete", server.WithSession(server.HandleError(server.NeedsAuthentication(server.DeleteImage)))).Methods(http.MethodPost)
	r.HandleFunc("/user/login", server.WithSession(server.HandleError(server.Login))).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", server.WithSession(server.HandleError(server.Logout))).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(server.NotFound)
	return r
}
