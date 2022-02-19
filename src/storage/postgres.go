package storage

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fredi12345/kuefa-karben/src/ent"
	"github.com/fredi12345/kuefa-karben/src/ent/user"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type PostgresBackend struct {
	client *ent.Client
}

func NewPostgresBackend() (*PostgresBackend, error) {
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	host := viper.GetString("postgres.host")
	port := viper.GetInt("postgres.port")
	database := viper.GetString("postgres.database")
	options := viper.GetStringSlice("postgres.options")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", user, password, host, port, database, strings.Join(options, "&"))

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not initialize database client: %w", err)
	}

	return &PostgresBackend{client: client}, nil
}

func (p *PostgresBackend) Migrate(ctx context.Context) error {
	return p.client.Schema.Create(ctx)
}

func (p *PostgresBackend) CreateEvent(event Event) (string, error) {
	e, err := p.client.Event.Create().
		SetTheme(event.Theme).
		SetStartingTime(event.EventDate).
		SetClosingTime(event.ClosingDate).
		SetStarter(event.Starter).
		SetMainDish(event.MainDish).
		SetDessert(event.Dessert).
		SetDescription(event.InfoText).
		SetTitleImage(event.ImageName).
		Save(context.TODO())
	if err != nil {
		return "", fmt.Errorf("could not create event: %w", err)
	}

	return e.ID.String(), nil
}

func (p *PostgresBackend) CreateParticipant(participant Participant) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) CreateComment(comment Comment) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) CreateImage(fileName string, event int) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetEvent(id int) (*Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetLatestEventId() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetComments(eventID int) ([]*Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetImages(eventId int) ([]*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetAllImages(page int, imagesPerSite int) ([]*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetImageCount() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetParticipants(eventId int) ([]*Participant, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetEventList(page int, eventsPerPage int) ([]*Event, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetEventCount() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetNewComments(limit int) ([]*Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) GetNewParticipants(limit int) ([]*Participant, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) UpdateEvent(event Event) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) DeleteComment(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) DeleteImage(id int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) DeleteParticipant(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) DeleteEvent(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresBackend) CreateUser(name, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password for user %s: %w", name, err)
	}

	err = p.client.User.Create().
		SetName(name).
		SetPassword(hash).
		Exec(context.TODO())
	if err != nil {
		if ent.IsConstraintError(err) {
			log.Printf("skip create user %s: %v\n", name, err)
			return nil
		}

		return fmt.Errorf("failed to create user %s: %w", name, err)
	}

	log.Printf("user %s created successfully\n", name)
	return nil
}

func (p *PostgresBackend) CheckCredentials(name, password string) (bool, error) {
	u, err := p.client.User.Query().
		Where(user.Name(name)).
		First(context.TODO())
	if err != nil {
		if ent.IsNotFound(err) {
			log.Printf("credential check for non-existing user %s failed\n", name)
			return false, nil
		}

		return false, fmt.Errorf("could not read user %s from database: %w", name, err)
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		return false, fmt.Errorf("error comparing password for user %s: %w", name, err)
	}

	return true, nil
}
