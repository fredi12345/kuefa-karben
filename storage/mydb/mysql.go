package mydb

import (
	"github.com/fredi12345/kuefa-karben/storage"
	"image"
)

type db struct {
}

func New() storage.Service {
	return &db{}
}

func (db *db) CreateEvent(event storage.Event) error {
	panic("implement me")
}

func (db *db) CreateParticipant(participant storage.Participant) error {
	panic("implement me")
}

func (db *db) CreateComment(comment storage.Comment) error {
	panic("implement me")
}

func (db *db) CreateImage(img image.Image, event int) error {
	panic("implement me")
}

func (db *db) CreateUser(name, password string) error {
	panic("implement me")
}

func (db *db) CheckCredentials(name, password string) (bool, error) {
	panic("implement me")
}
