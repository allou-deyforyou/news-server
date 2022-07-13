package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NewsTvSource holds the schema definition for the NewsTvSource entity.
type NewsTvSource struct {
	ent.Schema
}

// Fields of the NewsTvSource.
func (NewsTvSource) Fields() []ent.Field {
	return []ent.Field{
		field.String("logo").NotEmpty(),
		field.String("video").NotEmpty(),
		field.Bool("live").Default(true),
		field.Bool("status").Default(true),
		field.String("country").Default("ci"),
		field.String("description").NotEmpty(),
		field.String("language").Default("fr"),
		field.String("source").Unique().NotEmpty(),
		field.JSON("categories", map[string]string{}).Optional(),
	}
}

// Edges of the NewsTvSource.
func (NewsTvSource) Edges() []ent.Edge {
	return nil
}
