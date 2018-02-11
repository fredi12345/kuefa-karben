package storage

import (
	"image"
	"time"
)

type Service interface {
	CreateEvent(event Event) error
	CreateParticipant(participant Participant) error
	CreateComment(comment Comment) error
	CreateImage(img Image, event int) error
	CreateUser(name, password string) error

	GetEvent(id int) (*Event, error)
	GetComments(eventID int) ([]*Comment, error)
	GetImages(eventId int) ([]*Image, error)
	GetParticipants(eventId int) ([]*Participant, error)

	CheckCredentials(name, password string) (bool, error)
}

type Comment struct {
	Id      int
	Content string
	Name    string
	Created time.Time
	EventId int
}
type Event struct {
	Id        int
	Theme     string
	EventDate time.Time
	Created   time.Time
	Starter   string
	MainDish  string
	Dessert   string
	InfoText  string
	Img       image.Image
}

type Image struct {
	Id      int
	EventId int
	Picture image.Image
}

type Participant struct {
	Id      int
	Name    string
	Created time.Time
	Menu    int
	EventId int
}

type User struct {
	Id           int
	Name         string
	Salt         string
	PasswordHash string
}
