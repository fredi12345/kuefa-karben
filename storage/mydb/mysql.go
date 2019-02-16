package mydb

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"

	"github.com/fredi12345/kuefa-karben/random"
	"github.com/fredi12345/kuefa-karben/storage"
	"github.com/go-sql-driver/mysql"
)

const (
	dbCreateUser        = `INSERT INTO user ( name, salt, password) VALUES (?,?,?);`
	dbCreateEvent       = `INSERT INTO event (theme, event_date, starter, main_dish, dessert, infotext, image_name, created_date) VALUES (?,?,?,?,?,?,?, NOW())`
	dbCreateParticipant = `INSERT INTO participant (name, menu, message, event_id, participant_created) VALUES (?, ?, ?, ?, Now()) `
	dbCreateComment     = `INSERT INTO comment (content, name, comment_created, event_id) VALUES (?,?, Now(), ?)`
	dbCreateImage       = `INSERT INTO images (event_id, image_name) VALUES (?, ?)`

	dbGetEvent         = `SELECT event_id, theme, event_date, created_date, starter, main_dish, dessert, infotext, image_name FROM event WHERE event_id=?;`
	dbGetComments      = `SELECT comment.id, name, content, comment_created FROM comment WHERE event_id=? ORDER BY comment_created;`
	dbGetParticipants  = `SELECT participant.id, name, menu, message, participant_created, event_id FROM participant WHERE event_id=? ORDER BY participant_created;`
	dbGetImages        = `SELECT images.id, image_name FROM images WHERE event_id=? ORDER BY id`
	dbGetAllImages     = `SELECT images.id, images.image_name, e.event_id, theme FROM images INNER JOIN event e on images.event_id = e.event_id ORDER BY images.id DESC LIMIT ?,16`
	dbGetImageCount    = `SELECT COUNT(images.id) FROM images`
	dbGetSingleImage   = `SELECT image_name FROM images WHERE id=?`
	dbGetCredentials   = `SELECT salt, password FROM user WHERE name=?`
	dbGetLatestEventId = `SELECT event_id FROM event ORDER BY event_date DESC LIMIT 1`
	dbGetEventList     = `SELECT event_id,theme,event_date,image_name FROM event ORDER BY event_date DESC LIMIT ?,9 `
	dbGetEventCount    = `SELECT COUNT(event_id) FROM event`

	dbUpdateEvent = `UPDATE event SET theme=?, event_date=?, starter=?, main_dish=?, dessert=?, infotext=?, image_name=? WHERE event_id=?`

	dbDeleteComment     = `DELETE FROM comment WHERE id=?`
	dbDeleteImage       = `DELETE FROM images WHERE id=?`
	dbDeleteParticipant = `DELETE FROM participant WHERE id=?`
	dbDeleteEvent       = `DELETE FROM event WHERE event_id=?`
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

func (c *connection) DeleteImage(id int) (string, error) {
	var name string
	err := c.db.QueryRow(dbGetSingleImage, id).Scan(&name)
	if err != nil {
		return "", errors.WithStack(err)
	}

	_, err = c.db.Exec(dbDeleteImage, id)
	return name, errors.WithStack(err)
}

func (c *connection) DeleteComment(id int) error {
	_, err := c.db.Exec(dbDeleteComment, id)
	return errors.WithStack(err)
}

func (c *connection) DeleteParticipant(id int) error {
	_, err := c.db.Exec(dbDeleteParticipant, id)
	return errors.WithStack(err)
}

func (c *connection) DeleteEvent(id int) error {
	_, err := c.db.Exec(dbDeleteEvent, id)
	return errors.WithStack(err)
}

func (c *connection) GetEventList(page int) ([]*storage.Event, error) {
	var events []*storage.Event
	var offset = (page - 1) * 9
	rows, err := c.db.Query(dbGetEventList, offset)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Event
		err := rows.Scan(&resultItem.Id, &resultItem.Theme, &resultItem.EventDate, &resultItem.ImageName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		events = append(events, &resultItem)
	}

	return events, nil
}

//Event count to hide 'next page button' on last page
func (c *connection) GetEventCount() (int, error) {
	var count int
	row := c.db.QueryRow(dbGetEventCount)
	err := row.Scan(&count)
	return count, errors.WithStack(err)
}

func (c *connection) GetLatestEventId() (int, error) {
	var id int
	err := c.db.QueryRow(dbGetLatestEventId).Scan(&id)
	return id, errors.WithStack(err)
}

func (c *connection) GetEvent(id int) (*storage.Event, error) {
	event := storage.Event{}
	event.Id = id
	err := c.db.QueryRow(dbGetEvent, id).Scan(&event.Id, &event.Theme, &event.EventDate, &event.Created, &event.Starter, &event.MainDish, &event.Dessert, &event.InfoText, &event.ImageName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &event, nil
}

func (c *connection) GetComments(eventId int) ([]*storage.Comment, error) {
	var comments []*storage.Comment
	rows, err := c.db.Query(dbGetComments, eventId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Comment
		err := rows.Scan(&resultItem.Id, &resultItem.Name, &resultItem.Content, &resultItem.Created)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		comments = append(comments, &resultItem)
	}

	return comments, nil
}

func (c *connection) GetImages(eventId int) ([]*storage.Image, error) {
	var images []*storage.Image
	rows, err := c.db.Query(dbGetImages, eventId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Image
		err := rows.Scan(&resultItem.Id, &resultItem.Name)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		images = append(images, &resultItem)
	}

	return images, nil
}

func (c *connection) GetAllImages(page int) ([]*storage.Image, error) {
	var images []*storage.Image
	var offset = (page - 1) * 9
	rows, err := c.db.Query(dbGetAllImages, offset)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Image
		err := rows.Scan(&resultItem.Id, &resultItem.Name, &resultItem.EventId, &resultItem.EventName)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		images = append(images, &resultItem)
	}

	return images, nil
}

//Event count to hide 'next page button' on last page
func (c *connection) GetImageCount() (int, error) {
	var count int
	row := c.db.QueryRow(dbGetImageCount)
	err := row.Scan(&count)
	return count, errors.WithStack(err)
}

func (c *connection) GetParticipants(eventId int) ([]*storage.Participant, error) {
	var participants []*storage.Participant
	rows, err := c.db.Query(dbGetParticipants, eventId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultItem storage.Participant
		err := rows.Scan(&resultItem.Id, &resultItem.Name, &resultItem.Menu, &resultItem.Message, &resultItem.Created, &resultItem.EventId)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		participants = append(participants, &resultItem)
	}

	return participants, nil
}

func New(cfg *mysql.Config) (storage.Service, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &connection{db: db, rnd: random.New(time.Now().Unix())}, nil
}

func (c *connection) CreateEvent(event storage.Event) (int, error) {
	res, err := c.db.Exec(dbCreateEvent, event.Theme, event.EventDate, event.Starter, event.MainDish, event.Dessert, event.InfoText, event.ImageName)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return -1, errors.WithStack(ErrInputToLong)
		}
		return -1, errors.WithStack(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.WithStack(err)
	}
	return int(id), nil
}

func (c *connection) CreateParticipant(participant storage.Participant) error {
	_, err := c.db.Exec(dbCreateParticipant, participant.Name, participant.Menu, participant.Message, participant.EventId)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return errors.WithStack(ErrInputToLong)
		}
		return errors.WithStack(err)
	}
	return err
}

func (c *connection) CreateComment(comment storage.Comment) error {
	_, err := c.db.Exec(dbCreateComment, comment.Content, comment.Name, comment.EventId)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return errors.WithStack(ErrInputToLong)
		}
		return errors.WithStack(err)
	}
	return err
}

func (c *connection) CreateImage(name string, event int) error {
	_, err := c.db.Exec(dbCreateImage, event, name)
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1406 {
			return errors.WithStack(ErrInputToLong)
		}
		return errors.WithStack(err)
	}
	return err
}

func (c *connection) CreateUser(name, password string) error {
	salt := c.rnd.String(10)
	_, err := c.db.Exec(dbCreateUser, name, salt, hash(password, salt))
	if msqlErr, ok := err.(*mysql.MySQLError); ok {
		if msqlErr.Number == 1062 {
			return errors.WithStack(ErrUserAlreadyAssigned)
		}
		if msqlErr.Number == 1406 {
			return errors.WithStack(ErrUserToLong)
		}
		return errors.WithStack(err)
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
		return false, errors.WithStack(err)
	}

	hashedAttempt := hash(attemptedPassword, salt)
	return hashedAttempt == hashedPassword, nil
}

func (c *connection) UpdateEvent(e storage.Event) error {
	_, err := c.db.Exec(dbUpdateEvent, e.Theme, e.EventDate, e.Starter, e.MainDish, e.Dessert, e.InfoText, e.ImageName, e.Id)
	return errors.WithStack(err)
}

func hash(password string, salt string) string {
	hasher := sha256.New()
	io.WriteString(hasher, password)
	io.WriteString(hasher, salt)
	return hex.EncodeToString(hasher.Sum(nil))
}
