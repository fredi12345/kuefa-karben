package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New),
		field.Time("created").Default(time.Now).Immutable(),
		field.Time("last_modified").Default(time.Now).UpdateDefault(time.Now),
		field.String("theme").MaxLen(256),
		field.String("title_image"),
		field.Time("starting_time"),
		field.String("starter").MaxLen(512),
		field.String("main_dish").MaxLen(512),
		field.String("dessert").MaxLen(512),
		field.String("description").MaxLen(2048),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participants", Participant.Type),
		edge.To("comments", Comment.Type),
		edge.To("images", Image.Type),
	}
}
