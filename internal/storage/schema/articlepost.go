package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ArticlePost holds the schema definition for the ArticlePost entity.
type ArticlePost struct {
	ent.Schema
}

// Fields of the ArticlePostSchema.
func (ArticlePost) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("status").Default(true),
		field.String("title"),
		field.String("image"),
		field.Time("date").Optional(),
		field.String("link"),
		field.String("content").Optional(),
		field.String("source"),
	}
}

// Edges of the ArticlePost.
func (ArticlePost) Edges() []ent.Edge {
	return nil
}
