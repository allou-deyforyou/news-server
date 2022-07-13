package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NewsCategories holds the schema definition for the NewsCategories entity.
type NewsCategories struct {
	ent.Schema
}

// Fields of the NewsCategories.
func (NewsCategories) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("tv_categories"),
		field.Bool("status").Default(true),
		field.Strings("article_categories"),
		field.String("language").Default("fr"),
	}
}

// Edges of the NewsCategories.
func (NewsCategories) Edges() []ent.Edge {
	return nil
}
