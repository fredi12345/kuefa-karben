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
	dbName, user, password, err := config.Read("config.xml")
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	db, err := mydb.New(dbName, user, password)
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
	r.HandleFunc("/", server.Index).Methods(http.MethodGet)
	r.HandleFunc("/event/create", server.NeedsAuthentication(server.CreateEventPage)).Methods(http.MethodGet)
	r.HandleFunc("/event/edit", server.NeedsAuthentication(server.EditEventPage)).Methods(http.MethodGet)
	r.HandleFunc("/event/{id:[0-9]+}", server.EventDetail).Methods(http.MethodGet)
	r.HandleFunc("/event/all", server.AllEvents).Methods(http.MethodGet)
	r.HandleFunc("/impressum", server.Impressum).Methods(http.MethodGet)

	// http post methods
	r.HandleFunc("/event/add", server.NeedsAuthentication(server.AddEvent)).Methods(http.MethodPost)
	r.HandleFunc("/event/edit", server.NeedsAuthentication(server.EditEvent)).Methods(http.MethodPost)
	r.HandleFunc("/event/delete", server.NeedsAuthentication(server.DeleteEvent)).Methods(http.MethodPost)
	r.HandleFunc("/participant/add", server.AddParticipant).Methods(http.MethodPost)
	r.HandleFunc("/participant/delete", server.NeedsAuthentication(server.DeleteParticipant)).Methods(http.MethodPost)
	r.HandleFunc("/comment/add", server.AddComment).Methods(http.MethodPost)
	r.HandleFunc("/comment/delete", server.NeedsAuthentication(server.DeleteComment)).Methods(http.MethodPost)
	r.HandleFunc("/image/add", server.NeedsAuthentication(server.AddImage)).Methods(http.MethodPost)
	r.HandleFunc("/image/delete", server.NeedsAuthentication(server.DeleteImage)).Methods(http.MethodPost)
	r.HandleFunc("/user/login", server.Login).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", server.Logout).Methods(http.MethodPost)
	return r
}
