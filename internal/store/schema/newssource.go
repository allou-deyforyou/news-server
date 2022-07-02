package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NewsSource holds the schema definition for the NewsSource entity.
type NewsSource struct {
	ent.Schema
}

// Fields of the NewsSource.
func (NewsSource) Fields() []ent.Field {
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

		field.String("language").Default("fr"),
		field.String("country").Default("ci"),
		field.Bool("status").Default(true),
		field.String("name").Unique(),
		field.Strings("categories"),
		field.String("url"),
	}
}

// Edges of the NewsSource.
func (NewsSource) Edges() []ent.Edge {
	return nil
}
