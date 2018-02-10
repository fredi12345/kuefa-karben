package mydb

import (
	"database/sql"
	"fmt"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/go-sql-driver/mysql"
	"image"
)

const (
	dbCreateUser = `INSERT INTO user ( name, salt, password) VALUES (?,?,?)`
)

var (
	ErrUserAlreadyAssigned = fmt.Errorf("username already assigned")
	ErrUserToLong          = fmt.Errorf("username must be less than 256 characters")
	//ErrUserNotFound        = fmt.Errorf("user not found")
)

type connection struct {
	db *sql.DB
}

func New() storage.Service {
	return &connection{}
}

func (c *connection) CreateEvent(event storage.Event) error {
	panic("implement me")
}

func (c *connection) CreateParticipant(participant storage.Participant) error {
	panic("implement me")
}

func (c *connection) CreateComment(comment storage.Comment) error {
	panic("implement me")
}

func (c *connection) CreateImage(img image.Image, event int) error {
	panic("implement me")
}

func (c *connection) CreateUser(name, password string) error {
	salt := "12"         //TODO
	passwordHash := "12" //TODO
	_, err := c.db.Exec(dbCreateUser, name, salt, passwordHash)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1062 {
			return ErrUserAlreadyAssigned
		}
		if msqlErr.Number == 1406 {
			return ErrUserToLong
		}
		return fmt.Errorf("cannot execute statement: %v", err)
	}
	return err
}

func (c *connection) CheckCredentials(name, password string) (bool, error) {
	panic("implement me")
}
