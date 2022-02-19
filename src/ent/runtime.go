// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/fredi12345/kuefa-karben/src/ent/comment"
	"github.com/fredi12345/kuefa-karben/src/ent/event"
	"github.com/fredi12345/kuefa-karben/src/ent/image"
	"github.com/fredi12345/kuefa-karben/src/ent/participant"
	"github.com/fredi12345/kuefa-karben/src/ent/schema"
	"github.com/fredi12345/kuefa-karben/src/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreated is the schema descriptor for created field.
	commentDescCreated := commentFields[1].Descriptor()
	// comment.DefaultCreated holds the default value on creation for the created field.
	comment.DefaultCreated = commentDescCreated.Default.(func() time.Time)
	// commentDescName is the schema descriptor for name field.
	commentDescName := commentFields[2].Descriptor()
	// comment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	comment.NameValidator = commentDescName.Validators[0].(func(string) error)
	// commentDescMessage is the schema descriptor for message field.
	commentDescMessage := commentFields[3].Descriptor()
	// comment.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	comment.MessageValidator = commentDescMessage.Validators[0].(func(string) error)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.DefaultID holds the default value on creation for the id field.
	comment.DefaultID = commentDescID.Default.(func() uuid.UUID)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreated is the schema descriptor for created field.
	eventDescCreated := eventFields[1].Descriptor()
	// event.DefaultCreated holds the default value on creation for the created field.
	event.DefaultCreated = eventDescCreated.Default.(func() time.Time)
	// eventDescLastModified is the schema descriptor for last_modified field.
	eventDescLastModified := eventFields[2].Descriptor()
	// event.DefaultLastModified holds the default value on creation for the last_modified field.
	event.DefaultLastModified = eventDescLastModified.Default.(func() time.Time)
	// event.UpdateDefaultLastModified holds the default value on update for the last_modified field.
	event.UpdateDefaultLastModified = eventDescLastModified.UpdateDefault.(func() time.Time)
	// eventDescTheme is the schema descriptor for theme field.
	eventDescTheme := eventFields[3].Descriptor()
	// event.ThemeValidator is a validator for the "theme" field. It is called by the builders before save.
	event.ThemeValidator = eventDescTheme.Validators[0].(func(string) error)
	// eventDescStarter is the schema descriptor for starter field.
	eventDescStarter := eventFields[6].Descriptor()
	// event.StarterValidator is a validator for the "starter" field. It is called by the builders before save.
	event.StarterValidator = eventDescStarter.Validators[0].(func(string) error)
	// eventDescMainDish is the schema descriptor for main_dish field.
	eventDescMainDish := eventFields[7].Descriptor()
	// event.MainDishValidator is a validator for the "main_dish" field. It is called by the builders before save.
	event.MainDishValidator = eventDescMainDish.Validators[0].(func(string) error)
	// eventDescDessert is the schema descriptor for dessert field.
	eventDescDessert := eventFields[8].Descriptor()
	// event.DessertValidator is a validator for the "dessert" field. It is called by the builders before save.
	event.DessertValidator = eventDescDessert.Validators[0].(func(string) error)
	// eventDescDescription is the schema descriptor for description field.
	eventDescDescription := eventFields[9].Descriptor()
	// event.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	event.DescriptionValidator = eventDescDescription.Validators[0].(func(string) error)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() uuid.UUID)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescCreated is the schema descriptor for created field.
	imageDescCreated := imageFields[1].Descriptor()
	// image.DefaultCreated holds the default value on creation for the created field.
	image.DefaultCreated = imageDescCreated.Default.(func() time.Time)
	// imageDescFileName is the schema descriptor for file_name field.
	imageDescFileName := imageFields[2].Descriptor()
	// image.FileNameValidator is a validator for the "file_name" field. It is called by the builders before save.
	image.FileNameValidator = imageDescFileName.Validators[0].(func(string) error)
	// imageDescID is the schema descriptor for id field.
	imageDescID := imageFields[0].Descriptor()
	// image.DefaultID holds the default value on creation for the id field.
	image.DefaultID = imageDescID.Default.(func() uuid.UUID)
	participantFields := schema.Participant{}.Fields()
	_ = participantFields
	// participantDescCreated is the schema descriptor for created field.
	participantDescCreated := participantFields[1].Descriptor()
	// participant.DefaultCreated holds the default value on creation for the created field.
	participant.DefaultCreated = participantDescCreated.Default.(func() time.Time)
	// participantDescName is the schema descriptor for name field.
	participantDescName := participantFields[2].Descriptor()
	// participant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	participant.NameValidator = participantDescName.Validators[0].(func(string) error)
	// participantDescMessage is the schema descriptor for message field.
	participantDescMessage := participantFields[3].Descriptor()
	// participant.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	participant.MessageValidator = participantDescMessage.Validators[0].(func(string) error)
	// participantDescID is the schema descriptor for id field.
	participantDescID := participantFields[0].Descriptor()
	// participant.DefaultID holds the default value on creation for the id field.
	participant.DefaultID = participantDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreated is the schema descriptor for created field.
	userDescCreated := userFields[1].Descriptor()
	// user.DefaultCreated holds the default value on creation for the created field.
	user.DefaultCreated = userDescCreated.Default.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
