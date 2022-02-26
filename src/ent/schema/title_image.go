package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TitleImage holds the schema definition for the TitleImage entity.
type TitleImage struct {
	ent.Schema
}

// Fields of the TitleImage.
func (TitleImage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New),
		field.Time("created").Default(time.Now).Immutable(),
	}
}

// Edges of the TitleImage.
func (TitleImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("title_image").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			Unique(),
	}
}
