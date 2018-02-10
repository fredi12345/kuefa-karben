package storage

import (
	"time"
	"image"
)

type Service interface {
	CreateEvent(theme, starter, mainDish, dessert, infotext string, date time.Time, img image.Image) error
	CreateParticipant(name string, event, menu int) error
}
