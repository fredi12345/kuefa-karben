// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/fredi12345/kuefa-karben/src/ent/event"
	"github.com/google/uuid"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// LastModified holds the value of the "last_modified" field.
	LastModified time.Time `json:"last_modified,omitempty"`
	// Theme holds the value of the "theme" field.
	Theme string `json:"theme,omitempty"`
	// TitleImage holds the value of the "title_image" field.
	TitleImage string `json:"title_image,omitempty"`
	// StartingTime holds the value of the "starting_time" field.
	StartingTime time.Time `json:"starting_time,omitempty"`
	// ClosingTime holds the value of the "closing_time" field.
	ClosingTime *time.Time `json:"closing_time,omitempty"`
	// Starter holds the value of the "starter" field.
	Starter string `json:"starter,omitempty"`
	// MainDish holds the value of the "main_dish" field.
	MainDish string `json:"main_dish,omitempty"`
	// Dessert holds the value of the "dessert" field.
	Dessert string `json:"dessert,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges EventEdges `json:"edges"`
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Participants holds the value of the participants edge.
	Participants []*Participant `json:"participants,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Images holds the value of the images edge.
	Images []*Image `json:"images,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) ParticipantsOrErr() ([]*Participant, error) {
	if e.loadedTypes[0] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) ImagesOrErr() ([]*Image, error) {
	if e.loadedTypes[2] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldTheme, event.FieldTitleImage, event.FieldStarter, event.FieldMainDish, event.FieldDessert, event.FieldDescription:
			values[i] = new(sql.NullString)
		case event.FieldCreated, event.FieldLastModified, event.FieldStartingTime, event.FieldClosingTime:
			values[i] = new(sql.NullTime)
		case event.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case event.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				e.Created = value.Time
			}
		case event.FieldLastModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_modified", values[i])
			} else if value.Valid {
				e.LastModified = value.Time
			}
		case event.FieldTheme:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field theme", values[i])
			} else if value.Valid {
				e.Theme = value.String
			}
		case event.FieldTitleImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_image", values[i])
			} else if value.Valid {
				e.TitleImage = value.String
			}
		case event.FieldStartingTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field starting_time", values[i])
			} else if value.Valid {
				e.StartingTime = value.Time
			}
		case event.FieldClosingTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closing_time", values[i])
			} else if value.Valid {
				e.ClosingTime = new(time.Time)
				*e.ClosingTime = value.Time
			}
		case event.FieldStarter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field starter", values[i])
			} else if value.Valid {
				e.Starter = value.String
			}
		case event.FieldMainDish:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field main_dish", values[i])
			} else if value.Valid {
				e.MainDish = value.String
			}
		case event.FieldDessert:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dessert", values[i])
			} else if value.Valid {
				e.Dessert = value.String
			}
		case event.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		}
	}
	return nil
}

// QueryParticipants queries the "participants" edge of the Event entity.
func (e *Event) QueryParticipants() *ParticipantQuery {
	return (&EventClient{config: e.config}).QueryParticipants(e)
}

// QueryComments queries the "comments" edge of the Event entity.
func (e *Event) QueryComments() *CommentQuery {
	return (&EventClient{config: e.config}).QueryComments(e)
}

// QueryImages queries the "images" edge of the Event entity.
func (e *Event) QueryImages() *ImageQuery {
	return (&EventClient{config: e.config}).QueryImages(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return (&EventClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", created=")
	builder.WriteString(e.Created.Format(time.ANSIC))
	builder.WriteString(", last_modified=")
	builder.WriteString(e.LastModified.Format(time.ANSIC))
	builder.WriteString(", theme=")
	builder.WriteString(e.Theme)
	builder.WriteString(", title_image=")
	builder.WriteString(e.TitleImage)
	builder.WriteString(", starting_time=")
	builder.WriteString(e.StartingTime.Format(time.ANSIC))
	if v := e.ClosingTime; v != nil {
		builder.WriteString(", closing_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", starter=")
	builder.WriteString(e.Starter)
	builder.WriteString(", main_dish=")
	builder.WriteString(e.MainDish)
	builder.WriteString(", dessert=")
	builder.WriteString(e.Dessert)
	builder.WriteString(", description=")
	builder.WriteString(e.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

func (e Events) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
