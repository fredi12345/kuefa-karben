package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New),
		field.Time("created").Default(time.Now).Immutable(),
		field.String("name").MaxLen(256).Immutable(),
		field.String("message").MaxLen(1024).Immutable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("comments").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			Required().Unique(),
	}

}
