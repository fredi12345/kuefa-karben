package main

import (
	"github.com/fredi12345/kuefa-karben/config"
	"github.com/fredi12345/kuefa-karben/storage/mydb"
	"log"
	"net/http"
	"github.com/fredi12345/kuefa-karben/web"
)

func main() {
	dbName, user, password, err := config.Read("config.xml")
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	_, err = mydb.New(dbName, user, password)
	if err != nil {
		log.Fatalf("could not create database: %v", err)
	}

	fs := http.FileServer(http.Dir("resources/static"))
	http.Handle("/static/", fs)
	http.HandleFunc("/", web.Index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}


	// db.CreateUser("test", "12345")
}
