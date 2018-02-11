package storage

import (
	"image"
	"time"
)

type Service interface {
	CreateEvent(event Event) error
	CreateParticipant(participant Participant) error
	CreateComment(comment Comment) error
	CreateImage(img image.Image, event int) error
	CreateUser(name, password string) error

	GetEvent(id int) (*Event, error)
	GetComment(eventId int) ([]Comment, error)
	GetImages(eventId int) ([]image.Image, error)
	GetParticipants(eventId int) ([]Participant, error)

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
