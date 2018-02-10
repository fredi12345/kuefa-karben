package main

import (
	"github.com/fredi12345/kuefa-karben/config"
	"github.com/fredi12345/kuefa-karben/storage/mydb"
	"log"
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

}
