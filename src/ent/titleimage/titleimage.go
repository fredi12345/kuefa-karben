// Code generated by entc, DO NOT EDIT.

package titleimage

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the titleimage type in the database.
	Label = "title_image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the titleimage in the database.
	Table = "title_images"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "title_images"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_title_image"
)

// Columns holds all SQL columns for titleimage fields.
var Columns = []string{
	FieldID,
	FieldCreated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "title_images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_title_image",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
