package storage

import (
	"context"
	"time"
)

type Service interface {
	CreateEvent(event Event) (string, error)
	CreateParticipant(participant Participant) error
	CreateComment(comment Comment) error
	CreateImage(eventID string) (string, error)
	CreateTitleImage() (string, error)
	CreateUser(name, password string) error

	GetEvent(id string) (*Event, error)
	GetLatestEventId() (string, error)
	GetComments(eventID string) ([]*Comment, error)
	GetImages(eventId string) ([]*Image, error)
	GetAllImages(page int, imagesPerSite int) ([]*Image, error)
	GetImageCount() (int, error)
	GetParticipants(eventId string) ([]*Participant, error)
	GetEventList(offset int, limit int) ([]*Event, error)
	GetEventCount() (int, error)
	GetNewComments(limit int) ([]*Comment, error)
	GetNewParticipants(limit int) ([]*Participant, error)

	UpdateEvent(event Event) error

	DeleteComment(id string) error
	DeleteImage(id string) (string, error)
	DeleteParticipant(id string) error
	DeleteEvent(id string) error

	CheckCredentials(name, password string) (bool, error)
}

type Migrator interface {
	Migrate(ctx context.Context) error
}

type Comment struct {
	ID      string
	Content string
	Name    string
	Created time.Time
	EventID string
}

type Event struct {
	ID          string
	Theme       string
	EventDate   time.Time
	ClosingDate time.Time
	Created     time.Time
	Starter     string
	MainDish    string
	Dessert     string
	InfoText    string
	// Deprecated: use ImageID instead
	ImageName string
	ImageID   string
}

type Participant struct {
	ID              string
	Name            string
	Created         time.Time
	Message         string
	EventID         string
	ClassicCount    int
	VegetarianCount int
	VeganCount      int
}

type Image struct {
	ID        string
	Name      string
	EventID   string
	EventName string
}
