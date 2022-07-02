package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TvSource holds the schema definition for the TvSource entity.
type TvSource struct {
	ent.Schema
}

// Fields of the TvSource.
func (TvSource) Fields() []ent.Field {
	return []ent.Field{
		field.String("logo").NotEmpty(),
		field.String("video").NotEmpty(),
		field.String("title").NotEmpty(),
		field.Bool("status").Default(true),
		field.String("country").Default("ci"),
		field.String("description").NotEmpty(),
		field.String("language").Default("fr"),
	}
}

// Edges of the TvSource.
func (TvSource) Edges() []ent.Edge {
	return nil
}
