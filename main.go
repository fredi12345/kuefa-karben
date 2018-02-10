package main

import (
	"fmt"
	"github.com/fredi12345/kuefa-karben/config"
	"log"
)

func main() {
	db, user, pass, err := config.Read("config.xml")
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	fmt.Println(db)
	fmt.Println(user)
	fmt.Println(pass)
}
