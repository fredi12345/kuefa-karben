package mydb

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/SchiffFlieger/go-random"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/go-sql-driver/mysql"
	"io"
	"time"
)

const (
	dbCreateUser        = `INSERT INTO user ( name, salt, password) VALUES (?,?,?);`
	dbCreateEvent       = `INSERT INTO event (theme, event_date, created, starter, main_dish, dessert, infotext, image) VALUES (?,?, created = NOW(),?,?,?,?,?)`
	dbCreateParticipant = `INSERT INTO participant (name, created, menu, event_id) VALUES (?,created = Now(), ?, event_id = (SELECT id FROM Event ORDER BY  id LIMIT 1)) `
	dbCreateComment     = `INSERT INTO comment (content, name, created, event_id) VALUES (?,?, created=Now(),event_id = (SELECT id FROM Event ORDER BY  id LIMIT 1))`
	dbCreateImage       = `INSERT INTO images (event_id, picture) VALUES (event_id=(SELECT id FROM Event ORDER BY  id LIMIT 1), ?)`
	dbGetEvent          = `SELECT * FROM event WHERE id=?;`
	dbGetComments       = `SELECT * FROM comment WHERE event_id=? ORDER BY created;`
	dbGetParticipants   = `SELECT * FROM participant WHERE event_id=? ORDER BY created;`
	dbGetImages         = `SELECT * FROM images WHERE event_id=? ORDER BY id`
)

var (
	ErrUserAlreadyAssigned = fmt.Errorf("username already assigned")
	ErrUserToLong          = fmt.Errorf("username must be less than 256 characters")
	ErrInputToLong         = fmt.Errorf("some attribute was to long")
	//ErrUserNotFound        = fmt.Errorf("user not found")
)

type connection struct {
	db  *sql.DB
	rnd *random.Rnd
}

func (c *connection) GetEvent(id int) (*storage.Event, error) {
	event := storage.Event{}
	err := c.db.QueryRow(dbGetEvent, id).Scan(&event.Theme, &event.InfoText, &event.Starter, &event.MainDish, &event.Dessert, &event.Img, &event.Img, &event.Created, &event.EventDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("error scanning row: %v", err)
	}

	return &event, nil
}

func (c *connection) GetComments(eventId int) ([]*storage.Comment, error) {
	comments := []*storage.Comment{}
	rows, err := c.db.Query(dbGetComments, eventId)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Comment
		err := rows.Scan(&resultItem.Name, &resultItem.Content, &resultItem.Created, &resultItem.Id)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		comments = append(comments, &resultItem)
	}

	return comments, nil
}

func (c *connection) GetImages(eventId int) ([]*storage.Image, error) {
	images := []*storage.Image{}
	rows, err := c.db.Query(dbGetImages, eventId)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Image
		err := rows.Scan(&resultItem.Id, &resultItem.Picture)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		images = append(images, &resultItem)
	}

	return images, nil

}

func (c *connection) GetParticipants(eventId int) ([]*storage.Participant, error) {
	participants := []*storage.Participant{}
	rows, err := c.db.Query(dbGetParticipants, eventId)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Participant
		err := rows.Scan(&resultItem.Name, &resultItem.Menu, &resultItem.Id, &resultItem.Created)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		participants = append(participants, &resultItem)
	}

	return participants, nil
}

func New(dbName, user, password string) (storage.Service, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbName))
	if err != nil {
		return nil, fmt.Errorf("cannot open connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot ping database: %v", err)
	}

	return &connection{db: db, rnd: random.New(time.Now().Unix())}, nil
}

func (c *connection) CreateEvent(event storage.Event) error {
	_, err := c.db.Exec(dbCreateEvent, event.Theme, event.EventDate, event.Starter, event.MainDish, event.Dessert, event.InfoText, event.Img)
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

func (c *connection) CreateImage(img storage.Image, event int) error {
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
	salt := c.rnd.String(10)
	_, err := c.db.Exec(dbCreateUser, name, salt, hash(password, salt))
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

func hash(password string, salt string) string {
	hasher := sha256.New()
	io.WriteString(hasher, password)
	io.WriteString(hasher, salt)
	return hex.EncodeToString(hasher.Sum(nil))
}
