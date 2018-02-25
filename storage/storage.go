package storage

import (
	"time"
)

type Service interface {
	CreateEvent(event Event) error
	CreateParticipant(participant Participant) error
	CreateComment(comment Comment) error
	CreateImage(url string, event int) error
	CreateUser(name, password string) error

	GetEvent(id int) (*Event, error)
	GetLatestEventId() (int, error)
	GetComments(eventID int) ([]*Comment, error)
	GetImages(eventId int) ([]string, error)
	GetParticipants(eventId int) ([]*Participant, error)
	GetEventList() ([]*Event, error)

	DeleteComment(id int) error

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
	ImageUrl  string
}

type Participant struct {
	Name    string
	Created time.Time
	Menu    int
	EventId int
}
