package storage

import (
	"image"
	"time"
)

type Service interface {
	CreateEvent(theme, starter, mainDish, dessert, infotext string, date time.Time, img image.Image) error
	CreateParticipant(name string, event, menu int) error
	CreateComment(name, content string, event int) error
	CreateImage(img image.Image, event int) error
	CreateUser(name, password string) error

	CheckCredentials(name, password string) (bool, error)
}

type Event struct {
	Theme    string
	Starter  string
	MainDish string
	Dessert  string
	InfoText string
	Date     time.Time
	Img      image.Image
}

type Participant struct {
	Name string
	Menu int
}

type Comment struct {
	Content string
	Name    string
}

type User struct {
	Name     string
	Salt     string
	Password string
}
