package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		field.Time("starting_time"),
		field.Time("closing_time"),
		field.String("starter").MaxLen(512),
		field.String("main_dish").MaxLen(512),
		field.String("dessert").MaxLen(512),
		field.String("description").MaxLen(2048),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participants", Participant.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("comments", Comment.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("images", Image.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("title_image", TitleImage.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			Unique(),
	}
}
