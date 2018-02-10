package storage

import (
	"image"
	"time"
)

type Service interface {
	CreateEvent(theme, starter, mainDish, dessert, infotext string, date time.Time, img image.Image) error
	CreateParticipant(name string, event, menu int) error
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
