package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Categories holds the schema definition for the Categories entity.
type Categories struct {
	ent.Schema
}

// Fields of the CategoriesSchema.
func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("article_categories", map[string]string{}).Optional(),
		field.JSON("media_categories", map[string]string{}).Optional(),
		field.String("language").Default("fr"),
		field.Bool("status").Default(true),

	}
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return nil
}
