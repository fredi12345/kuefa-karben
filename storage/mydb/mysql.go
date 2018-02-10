package mydb

import (
	"database/sql"
	"fmt"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/go-sql-driver/mysql"
	"image"
)

const (
	dbCreateUser        = `INSERT INTO user ( name, salt, password) VALUES (?,?,?);`
	dbCreateEvent       = `INSERT INTO event (theme, event_date, created, starter, main_dish, dessert, infotext, image) VALUES (?,?, created = NOW(),?,?,?,?,?)`
	dbCreateParticipant = `INSERT INTO participant (name, created, menu, event_id) VALUES (?,created = Now(), ?, event_id = (SELECT id FROM Event ORDER BY  id LIMIT 1)) `
	dbCreateComment     = `INSERT INTO comment (content, name, created, event_id) VALUES (?,?, created=Now(),event_id = (SELECT id FROM Event ORDER BY  id LIMIT 1))`
	dbCreateImage       = `INSERT INTO images (event_id, picture) VALUES (event_id=(SELECT id FROM Event ORDER BY  id LIMIT 1), ?)`
)

var (
	ErrUserAlreadyAssigned = fmt.Errorf("username already assigned")
	ErrUserToLong          = fmt.Errorf("username must be less than 256 characters")
	ErrInputToLong         = fmt.Errorf("some attribute was to long")
	//ErrUserNotFound        = fmt.Errorf("user not found")
)

type connection struct {
	db *sql.DB
}

func New() storage.Service {
	return &connection{}
}

func (c *connection) CreateEvent(event storage.Event) error {
	_, err := c.db.Exec(dbCreateEvent, event.Theme, event.Date, event.Starter, event.MainDish, event.Dessert, event.InfoText, event.Img)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return ErrInputToLong
		}
		return fmt.Errorf("cannot execute statement: %v", err)
	}
	return err
}

func (c *connection) CreateParticipant(participant storage.Participant) error {
	_, err := c.db.Exec(dbCreateParticipant, participant.Name, participant.Menu)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return ErrInputToLong
		}
		return fmt.Errorf("cannot execute statement: %v", err)
	}
	return err
}

func (c *connection) CreateComment(comment storage.Comment) error {
	_, err := c.db.Exec(dbCreateComment, comment.Content, comment.Name)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return ErrInputToLong
		}
		return fmt.Errorf("cannot execute statement: %v", err)
	}
	return err
}

func (c *connection) CreateImage(img image.Image, event int) error {
	_, err := c.db.Exec(dbCreateImage, img)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return ErrInputToLong
		}
		return fmt.Errorf("cannot execute statement: %v", err)
	}
	return err
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
