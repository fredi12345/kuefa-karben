package mydb

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/SchiffFlieger/go-random"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/go-sql-driver/mysql"
)

const (
	dbCreateUser        = `INSERT INTO user ( name, salt, password) VALUES (?,?,?);`
	dbCreateEvent       = `INSERT INTO event (theme, event_date, starter, main_dish, dessert, infotext, image_url, created_date) VALUES (?,?,?,?,?,?,?, NOW())`
	dbCreateParticipant = `INSERT INTO participant (name, participant_created, menu, event_id) VALUES (?, Now(), ?, (SELECT event_id FROM Event ORDER BY  event_id LIMIT 1)) `
	dbCreateComment     = `INSERT INTO comment (content, name, comment_created, event_id) VALUES (?,?, Now(), (SELECT event_id FROM Event ORDER BY  event_id LIMIT 1))`
	dbCreateImage       = `INSERT INTO images (event_id, image_url) VALUES (?, ?)`

	dbGetEvent         = `SELECT theme, event_date, created_date, starter, main_dish, dessert, infotext FROM event WHERE event_id=?;`
	dbGetComments      = `SELECT name, content, comment_created FROM comment WHERE event_id=? ORDER BY comment_created;`
	dbGetParticipants  = `SELECT name, menu, participant_created, event_id FROM participant WHERE event_id=? ORDER BY participant_created;`
	dbGetImages        = `SELECT image_url FROM images WHERE event_id=? ORDER BY id`
	dbGetCredentials   = `SELECT salt, password FROM user WHERE name=?`
	dbGetLatestEventId = `SELECT event_id FROM event ORDER BY event_date DESC LIMIT 1`
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

func (c *connection) GetLatestEventId() (int, error) {
	var id int
	err := c.db.QueryRow(dbGetLatestEventId).Scan(&id)
	return id, err
}

func (c *connection) GetEvent(id int) (*storage.Event, error) {
	event := storage.Event{}
	event.Id = id
	err := c.db.QueryRow(dbGetEvent, id).Scan(&event.Theme, &event.EventDate, &event.Created, &event.Starter, &event.MainDish, &event.Dessert, &event.InfoText)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("mysql.go|GetEvent: error scanning row: %v", err)
	}
	return &event, nil
}

func (c *connection) GetComments(eventId int) ([]*storage.Comment, error) {
	var comments []*storage.Comment
	rows, err := c.db.Query(dbGetComments, eventId)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Comment
		err := rows.Scan(&resultItem.Name, &resultItem.Content, &resultItem.Created)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		comments = append(comments, &resultItem)
	}

	return comments, nil
}

func (c *connection) GetImages(eventId int) ([]string, error) {
	var urls []string
	rows, err := c.db.Query(dbGetImages, eventId)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem string
		err := rows.Scan(&resultItem)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		urls = append(urls, resultItem)
	}

	return urls, nil
}

func (c *connection) GetParticipants(eventId int) ([]*storage.Participant, error) {
	var participants []*storage.Participant
	rows, err := c.db.Query(dbGetParticipants, eventId)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Participant
		err := rows.Scan(&resultItem.Name, &resultItem.Menu, &resultItem.Created, &resultItem.EventId)
		if err != nil {
			return nil, fmt.Errorf("mysql.go|GetParticipants: error scanning row: %v", err)
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
	_, err := c.db.Exec(dbCreateEvent, event.Theme, event.EventDate, event.Starter, event.MainDish, event.Dessert, event.InfoText, event.ImageUrl)
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

func (c *connection) CreateImage(url string, event int) error {
	_, err := c.db.Exec(dbCreateImage, event, url)
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

func (c *connection) CheckCredentials(name, attemptedPassword string) (bool, error) {
	var salt, hashedPassword string
	err := c.db.QueryRow(dbGetCredentials, name).Scan(&salt, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error scanning row: %v", err)
	}

	hashedAttempt := hash(attemptedPassword, salt)
	return hashedAttempt == hashedPassword, nil
}

func hash(password string, salt string) string {
	hasher := sha256.New()
	io.WriteString(hasher, password)
	io.WriteString(hasher, salt)
	return hex.EncodeToString(hasher.Sum(nil))
}
