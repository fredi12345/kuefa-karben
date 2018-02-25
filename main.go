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
	r.HandleFunc("/", server.Index).Methods(http.MethodGet)
	r.HandleFunc("/impressum", server.Impressum).Methods(http.MethodGet)
	r.HandleFunc("/participate", server.Participate).Methods(http.MethodPost)
	r.HandleFunc("/upload", server.Upload).Methods(http.MethodPost)
	r.HandleFunc("/login", server.Login).Methods(http.MethodPost)
	r.HandleFunc("/create", server.Create).Methods(http.MethodPost)
	r.HandleFunc("/comment", server.Comment).Methods(http.MethodPost)
	r.HandleFunc("/delete/comment", server.DeleteComment).Methods(http.MethodPost)
	return r
}
