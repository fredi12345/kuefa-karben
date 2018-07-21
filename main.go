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

	server := web.NewServer(db, path.Join("resources", "public", "images"))
	handler := createHandler(server)
	fmt.Println("http://localhost:8080")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

func createHandler(server *web.Server) http.Handler {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("resources/public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	// http get methods
	r.HandleFunc("/", server.HandleError(server.WithSession(server.Index))).Methods(http.MethodGet)
	r.HandleFunc("/event/create", server.HandleError(server.WithSession(server.NeedsAuthentication(server.CreateEventPage)))).Methods(http.MethodGet)
	r.HandleFunc("/event/edit", server.HandleError(server.WithSession(server.NeedsAuthentication(server.EditEventPage)))).Methods(http.MethodGet)
	r.HandleFunc("/event/{id:[0-9]+}", server.HandleError(server.WithSession(server.EventDetail))).Methods(http.MethodGet)
	r.HandleFunc("/event/all", server.HandleError(server.WithSession(server.AllEvents))).Methods(http.MethodGet)
	r.HandleFunc("/impressum", server.HandleError(server.WithSession(server.Impressum))).Methods(http.MethodGet)

	// http post methods
	r.HandleFunc("/event/add", server.HandleError(server.WithSession(server.NeedsAuthentication(server.AddEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/event/edit", server.HandleError(server.WithSession(server.NeedsAuthentication(server.EditEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/event/delete", server.HandleError(server.WithSession(server.NeedsAuthentication(server.DeleteEvent)))).Methods(http.MethodPost)
	r.HandleFunc("/participant/add", server.HandleError(server.WithSession(server.AddParticipant))).Methods(http.MethodPost)
	r.HandleFunc("/participant/delete", server.HandleError(server.WithSession(server.NeedsAuthentication(server.DeleteParticipant)))).Methods(http.MethodPost)
	r.HandleFunc("/comment/add", server.HandleError(server.WithSession(server.AddComment))).Methods(http.MethodPost)
	r.HandleFunc("/comment/delete", server.HandleError(server.WithSession(server.NeedsAuthentication(server.DeleteComment)))).Methods(http.MethodPost)
	r.HandleFunc("/image/add", server.HandleError(server.WithSession(server.NeedsAuthentication(server.AddImage)))).Methods(http.MethodPost)
	r.HandleFunc("/image/delete", server.HandleError(server.WithSession(server.NeedsAuthentication(server.DeleteImage)))).Methods(http.MethodPost)
	r.HandleFunc("/user/login", server.HandleError(server.WithSession(server.Login))).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", server.HandleError(server.WithSession(server.Logout))).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(server.NotFound)
	return r
}
