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
		field.Bool("status").Default(true),
		field.String("language").Default("fr"),
		field.JSON("tv_categories", map[string]string{}),
		field.JSON("article_categories", map[string]string{}),
	}
}

// Edges of the NewsCategories.
func (NewsCategories) Edges() []ent.Edge {
	return nil
}
