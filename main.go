package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fredi12345/kuefa-karben/config"
	"github.com/fredi12345/kuefa-karben/storage/mydb"
	"github.com/fredi12345/kuefa-karben/web"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Read("config.xml")
	if err != nil {
		log.Fatalf("could not read config: %v\n", err)
	}

	db, err := mydb.New(cfg)
	if err != nil {
		log.Fatalf("could not create database: %v\n", err)
	}

	if err := os.MkdirAll(cfg.Path.Image, 0750|os.ModeDir); err != nil {
		log.Fatalf("could not create folder: %v\n", err)
	}

	if err := os.MkdirAll(cfg.Path.Thumbnail, 0750|os.ModeDir); err != nil {
		log.Fatalf("could not create folder: %v\n", err)
	}

	server, err := web.NewServer(db, cfg.Path.Image, cfg.Path.Thumbnail, "cookies.key")
	if err != nil {
		log.Fatalf("could not create server: %v\n", err)
	}

	handler := createHandler(server)
	fmt.Printf("http://localhost:%s\n", cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handler); err != nil {
		log.Fatal(err)
	}
}

func createHandler(server *web.Server) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.Dir("resources/public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public", blockDirectoryListing(fs, server)))
	r.Handle("/robots.txt", fs)

	// redirect from incomplete urls
	r.Handle("/event", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/event/all/0", http.RedirectHandler("/event/all/1", http.StatusSeeOther))
	r.Handle("/gallery", http.RedirectHandler("/gallery/1", http.StatusSeeOther))
	r.Handle("/gallery/0", http.RedirectHandler("/gallery/1", http.StatusSeeOther))

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

func blockDirectoryListing(next http.Handler, server *web.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			server.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
