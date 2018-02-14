package main

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/config"
	"github.com/fredi12345/kuefa-karben/storage/mydb"
	"github.com/fredi12345/kuefa-karben/web"
	"log"
	"net/http"
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

	server := web.NewServer(db)

	fs := http.FileServer(http.Dir("resources/public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", server.Index)
	http.HandleFunc("/participate", server.Participate)
	fmt.Println("http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// db.CreateUser("test", "12345")
}
