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
	GetImages(eventId int) ([]*Image, error)
	GetParticipants(eventId int) ([]*Participant, error)
	GetEventList() ([]*Event, error)

	UpdateEvent(event Event) error
	UpdateEventImage(id int, url string) error

	DeleteComment(id int) error
	DeleteImage(id int) (string, error)
	DeleteParticipant(id int) error
	DeleteEvent(id int) error

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
	Id      int
	Name    string
	Created time.Time
	Menu    int
	EventId int
}

var menuToString = map[int]string{0: "Normal", 1: "Vegetarisch", 2: "Vegan"}

func (p *Participant) MenuString() string {
	return menuToString[p.Menu]
}

type Image struct {
	Id  int
	URL string
}
