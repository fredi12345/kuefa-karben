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
	GetComments(eventID int) ([]*Comment, error)
	//GetImages(eventId int) ([]*image.Image, error)
	GetParticipants(eventId int) ([]*Participant, error)

	CheckCredentials(name, password string) (bool, error)
}

type Comment struct {
	Content string
	Name    string
	Created time.Time
	EventId int
}
type Event struct {
	Id           int
	Theme        string
	EventDate    time.Time
	Created      time.Time
	Starter      string
	MainDish     string
	Dessert      string
	InfoText     string
	Img          image.Image
	Participants []*Participant
}

type Participant struct {
	Name    string
	Created time.Time
	Menu    int
	EventId int
}
