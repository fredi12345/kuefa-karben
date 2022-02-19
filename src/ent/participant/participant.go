// Code generated by entc, DO NOT EDIT.

package participant

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the participant type in the database.
	Label = "participant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldMenu holds the string denoting the menu field in the database.
	FieldMenu = "menu"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the participant in the database.
	Table = "participants"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "participants"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_participants"
)

// Columns holds all SQL columns for participant fields.
var Columns = []string{
	FieldID,
	FieldCreated,
	FieldName,
	FieldMessage,
	FieldMenu,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "participants"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_participants",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Menu defines the type for the "menu" enum field.
type Menu string

// Menu values.
const (
	MenuCLASSIC    Menu = "CLASSIC"
	MenuVEGETARIAN Menu = "VEGETARIAN"
	MenuVEGAN      Menu = "VEGAN"
)

func (m Menu) String() string {
	return string(m)
}

// MenuValidator is a validator for the "menu" field enum values. It is called by the builders before save.
func MenuValidator(m Menu) error {
	switch m {
	case MenuCLASSIC, MenuVEGETARIAN, MenuVEGAN:
		return nil
	default:
		return fmt.Errorf("participant: invalid enum value for menu field: %q", m)
	}
}
