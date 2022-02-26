package storage

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fredi12345/kuefa-karben/src/ent"
	"github.com/fredi12345/kuefa-karben/src/ent/comment"
	"github.com/fredi12345/kuefa-karben/src/ent/event"
	"github.com/fredi12345/kuefa-karben/src/ent/image"
	"github.com/fredi12345/kuefa-karben/src/ent/participant"
	"github.com/fredi12345/kuefa-karben/src/ent/user"
	"github.com/google/uuid"
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
	imageUUID, _ := uuid.Parse(event.ImageID)
	e, err := p.client.Event.Create().
		SetTheme(event.Theme).
		SetStartingTime(event.EventDate).
		SetClosingTime(event.ClosingDate).
		SetStarter(event.Starter).
		SetMainDish(event.MainDish).
		SetDessert(event.Dessert).
		SetDescription(event.InfoText).
		SetTitleImageID(imageUUID).
		Save(context.TODO())
	if err != nil {
		return "", fmt.Errorf("could not create event: %w", err)
	}

	return e.ID.String(), nil
}

func (p *PostgresBackend) CreateParticipant(participant Participant) error {
	eventID, _ := uuid.Parse(participant.EventID)
	err := p.client.Participant.Create().
		SetName(participant.Name).
		SetClassicMenu(participant.ClassicCount).
		SetVegetarianMenu(participant.VegetarianCount).
		SetVeganMenu(participant.VeganCount).
		SetMessage(participant.Message).
		SetEventID(eventID).
		Exec(context.TODO())
	if err != nil {
		return fmt.Errorf("could not create participant for event %s: %w", eventID.String(), err)
	}

	return nil
}

func (p *PostgresBackend) CreateComment(comment Comment) error {
	eventID, _ := uuid.Parse(comment.EventID)
	err := p.client.Comment.Create().
		SetName(comment.Name).
		SetMessage(comment.Content).
		SetEventID(eventID).
		Exec(context.TODO())
	if err != nil {
		return fmt.Errorf("could not create comment for event %s: %w", eventID.String(), err)
	}

	return nil
}

func (p *PostgresBackend) CreateImage(eventID string) (string, error) {
	eventUUID, _ := uuid.Parse(eventID)
	img, err := p.client.Image.Create().
		SetEventID(eventUUID).
		Save(context.TODO())
	if err != nil {
		return "", fmt.Errorf("could not save image for event %s: %w", eventUUID.String(), err)
	}

	return img.ID.String(), nil
}

func (p *PostgresBackend) CreateTitleImage() (string, error) {
	img, err := p.client.TitleImage.Create().
		Save(context.TODO())
	if err != nil {
		return "", fmt.Errorf("could not save new image: %w", err)
	}

	return img.ID.String(), nil
}

func (p *PostgresBackend) GetEvent(id string) (*Event, error) {
	eventUUID, _ := uuid.Parse(id)

	ev, err := p.client.Event.Get(context.TODO(), eventUUID)
	if err != nil {
		return nil, fmt.Errorf("could not read event %s: %w", id, err)
	}

	return &Event{
		ID:          ev.ID.String(),
		Theme:       ev.Theme,
		Created:     ev.Created,
		EventDate:   ev.StartingTime,
		ClosingDate: ev.ClosingTime,
		Starter:     ev.Starter,
		MainDish:    ev.MainDish,
		Dessert:     ev.Dessert,
		InfoText:    ev.Description,
		ImageName:   ev.Edges.TitleImage.ID.String(),
	}, nil
}

func (p *PostgresBackend) GetLatestEventId() (string, error) {
	ev, err := p.client.Event.Query().
		Select(event.FieldID).
		Order(ent.Desc(event.FieldStartingTime)).
		First(context.TODO())
	if err != nil {
		return "", fmt.Errorf("could not determine latest eventID: %w", err)
	}

	return ev.ID.String(), nil
}

func (p *PostgresBackend) GetComments(eventID string) ([]*Comment, error) {
	eventUUID, _ := uuid.Parse(eventID)
	dbComments, err := p.client.Comment.Query().
		Where(comment.HasEventWith(event.ID(eventUUID))).
		Order(ent.Desc(comment.FieldCreated)).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read comments for event %s: %w", eventID, err)
	}

	comments := make([]*Comment, 0, len(dbComments))
	for _, c := range dbComments {
		comments = append(comments, &Comment{
			ID:      c.ID.String(),
			Content: c.Message,
			Name:    c.Name,
			Created: c.Created,
			EventID: c.Edges.Event.ID.String(),
		})
	}

	return comments, nil
}

func (p *PostgresBackend) GetImages(eventID string) ([]*Image, error) {
	eventUUID, _ := uuid.Parse(eventID)
	dbImages, err := p.client.Image.Query().
		Where(image.HasEventWith(event.ID(eventUUID))).
		Order(ent.Desc(image.FieldCreated)).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read images for event %s: %w", eventID, err)
	}

	images := make([]*Image, 0, len(dbImages))
	for _, i := range dbImages {
		images = append(images, &Image{
			ID:      i.ID.String(),
			Name:    i.ID.String(),
			EventID: i.Edges.Event.ID.String(),
		})
	}

	return images, nil
}

func (p *PostgresBackend) GetAllImages(page int, imagesPerSite int) ([]*Image, error) {
	offset := (page - 1) * imagesPerSite
	dbImages, err := p.client.Image.Query().WithEvent().
		Order(ent.Desc(image.FieldCreated)).
		Offset(offset).
		Limit(imagesPerSite).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read image page %d: %w", page, err)
	}

	images := make([]*Image, 0, len(dbImages))
	for _, i := range dbImages {
		images = append(images, &Image{
			ID:        i.ID.String(),
			Name:      i.ID.String(),
			EventID:   i.Edges.Event.ID.String(),
			EventName: i.Edges.Event.Theme,
		})
	}

	return images, nil
}

func (p *PostgresBackend) GetImageCount() (int, error) {
	count, err := p.client.Image.Query().Count(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("could not determine total image count: %w", err)
	}

	return count, nil
}

func (p *PostgresBackend) GetParticipants(eventID string) ([]*Participant, error) {
	eventUUID, _ := uuid.Parse(eventID)
	dbParticipants, err := p.client.Participant.Query().
		Where(participant.HasEventWith(event.ID(eventUUID))).
		Order(ent.Desc(image.FieldCreated)).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read participants for event %s: %w", eventID, err)
	}

	participants := make([]*Participant, 0, len(dbParticipants))
	for _, pa := range dbParticipants {
		participants = append(participants, &Participant{
			ID:              pa.ID.String(),
			Created:         pa.Created,
			EventID:         pa.Edges.Event.ID.String(),
			Name:            pa.Name,
			Message:         pa.Message,
			ClassicCount:    pa.ClassicMenu,
			VegetarianCount: pa.VegetarianMenu,
			VeganCount:      pa.VeganMenu,
		})
	}

	return participants, nil
}

func (p *PostgresBackend) GetEventList(page int, eventsPerPage int) ([]*Event, error) {
	offset := (page - 1) * eventsPerPage
	dbEvents, err := p.client.Event.Query().
		Order(ent.Desc(event.FieldStartingTime)).
		Offset(page).
		Limit(offset).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read event page %d: %w", page, err)
	}

	events := make([]*Event, 0, len(dbEvents))
	for _, e := range dbEvents {
		events = append(events, &Event{
			ID:        e.ID.String(),
			Theme:     e.Theme,
			Created:   e.Created,
			EventDate: e.StartingTime,
			ImageName: e.Edges.TitleImage.ID.String(),
		})
	}

	return events, nil
}

func (p *PostgresBackend) GetEventCount() (int, error) {
	count, err := p.client.Event.Query().Count(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("could not determine total event count: %w", err)
	}

	return count, nil
}

func (p *PostgresBackend) GetNewComments(limit int) ([]*Comment, error) {
	dbComments, err := p.client.Comment.Query().
		Order(ent.Desc(comment.FieldCreated)).
		Limit(limit).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read latest comments: %w", err)
	}

	comments := make([]*Comment, 0, len(dbComments))
	for _, c := range dbComments {
		comments = append(comments, &Comment{
			ID:      c.ID.String(),
			Content: c.Message,
			Name:    c.Name,
			Created: c.Created,
			EventID: c.Edges.Event.ID.String(),
		})
	}

	return comments, nil
}

func (p *PostgresBackend) GetNewParticipants(limit int) ([]*Participant, error) {
	dbParticipants, err := p.client.Participant.Query().
		Order(ent.Desc(image.FieldCreated)).
		Limit(limit).
		All(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("could not read latest participants: %w", err)
	}

	participants := make([]*Participant, 0, len(dbParticipants))
	for _, pa := range dbParticipants {
		participants = append(participants, &Participant{
			ID:              pa.ID.String(),
			Created:         pa.Created,
			EventID:         pa.Edges.Event.ID.String(),
			Name:            pa.Name,
			Message:         pa.Message,
			ClassicCount:    pa.ClassicMenu,
			VegetarianCount: pa.VegetarianMenu,
			VeganCount:      pa.VeganMenu,
		})
	}

	return participants, nil
}

func (p *PostgresBackend) UpdateEvent(event Event) error {
	eventUUID, _ := uuid.Parse(event.ID)
	imageUUID, _ := uuid.Parse(event.ImageID)
	_, err := p.client.Event.UpdateOneID(eventUUID).
		SetTheme(event.Theme).
		SetStartingTime(event.EventDate).
		SetClosingTime(event.ClosingDate).
		SetStarter(event.Starter).
		SetMainDish(event.MainDish).
		SetDessert(event.Dessert).
		SetDescription(event.InfoText).
		SetTitleImageID(imageUUID).
		Save(context.TODO())
	if err != nil {
		return fmt.Errorf("could not update event %s: %w", eventUUID.String(), err)
	}

	return nil
}

func (p *PostgresBackend) DeleteComment(id string) error {
	commentID, _ := uuid.Parse(id)
	err := p.client.Comment.DeleteOneID(commentID).Exec(context.TODO())
	if err != nil {
		return fmt.Errorf("could not delete comment %s: %w", id, err)
	}

	return nil
}

func (p *PostgresBackend) DeleteImage(id string) (string, error) {
	imageID, _ := uuid.Parse(id)
	img, err := p.client.Image.Get(context.TODO(), imageID)
	if err != nil {
		return "", fmt.Errorf("could not retrieve file name for image %s: %w", id, err)
	}

	err = p.client.Image.DeleteOneID(imageID).Exec(context.TODO())
	if err != nil {
		return "", fmt.Errorf("could not delete image %s: %w", id, err)
	}

	return img.ID.String(), nil
}

func (p *PostgresBackend) DeleteParticipant(id string) error {
	participantID, _ := uuid.Parse(id)
	err := p.client.Participant.DeleteOneID(participantID).Exec(context.TODO())
	if err != nil {
		return fmt.Errorf("could not delete participant %s: %w", id, err)
	}

	return nil
}

func (p *PostgresBackend) DeleteEvent(id string) error {
	eventID, _ := uuid.Parse(id)
	err := p.client.Event.DeleteOneID(eventID).Exec(context.TODO())
	if err != nil {
		return fmt.Errorf("could not delete event %s: %w", id, err)
	}

	return nil
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
		Only(context.TODO())
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
