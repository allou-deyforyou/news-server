package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NewsArticleSource holds the schema definition for the NewsArticleSource entity.
type NewsArticleSource struct {
	ent.Schema
}

// Fields of the NewsArticleSource.
func (NewsArticleSource) Fields() []ent.Field {
	return []ent.Field{
		field.String("latest_post_url").
			Optional().Nillable(),
		field.JSON("latest_post_selector", &NewsPostSelector{}).
			Optional(),

		field.String("category_post_url").
			Optional().Nillable(),
		field.JSON("category_post_selector", &NewsPostSelector{}).
			Optional(),

		field.JSON("article_selector", &NewsArticleSelector{}).
			Optional(),

		field.JSON("categories", map[string]string{}),
		field.String("language").Default("fr"),
		field.String("country").Default("ci"),
		field.Bool("status").Default(true),
		field.String("logo").NotEmpty(),
		field.String("name").Unique(),
		field.String("url"),
	}
}

// Edges of the NewsArticleSource.
func (NewsArticleSource) Edges() []ent.Edge {
	return nil
}
