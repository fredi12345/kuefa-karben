package storage

import (
	"time"
)

type Service interface {
	CreateEvent(event Event) (int, error)
	CreateParticipant(participant Participant) error
	CreateComment(comment Comment) error
	CreateImage(fileName string, event int) error
	CreateUser(name, password string) error

	GetEvent(id int) (*Event, error)
	GetLatestEventId() (int, error)
	GetComments(eventID int) ([]*Comment, error)
	GetImages(eventId int) ([]*Image, error)
	GetAllImages(page int, imagesPerSite int) ([]*Image, error)
	GetImageCount() (int, error)
	GetParticipants(eventId int) ([]*Participant, error)
	GetEventList(page int, eventsPerPage int) ([]*Event, error)
	GetEventCount() (int, error)
	GetNewComments(limit int) ([]*Comment, error)
	GetNewParticipants(limit int) ([]*Participant, error)

	UpdateEvent(event Event) error

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
	ImageName string
}

type Participant struct {
	Id              int
	Name            string
	Created         time.Time
	Message         string
	EventId         int
	ClassicCount    int
	VegetarianCount int
	VeganCount      int
}

type Image struct {
	Id        int
	Name      string
	EventId   int
	EventName string
}
