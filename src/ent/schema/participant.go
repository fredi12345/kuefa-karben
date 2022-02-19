package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Participant holds the schema definition for the Participant entity.
type Participant struct {
	ent.Schema
}

// Fields of the Participant.
func (Participant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New),
		field.Time("created").Default(time.Now).Immutable(),
		field.String("name").MaxLen(256).Immutable(),
		field.String("message").MaxLen(256).Immutable(),
		field.Int("classic_menu").Min(0).Immutable(),
		field.Int("vegetarian_menu").Min(0).Immutable(),
		field.Int("vegan_menu").Min(0).Immutable(),
	}
}

// Edges of the Participant.
func (Participant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("participants").Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}

}
