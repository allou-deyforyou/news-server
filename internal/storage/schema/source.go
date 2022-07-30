package schema

import (
	"news/internal/storage/custom"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Source holds the schema definition for the Source entity.
type Source struct {
	ent.Schema
}

// Fields of the Source.
func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("article_featured_post_selector", &custom.SourcePostSelector{}).Optional(),
		field.String("article_featured_post_url").Optional(),
		field.JSON("article_category_post_selector", &custom.SourcePostSelector{}).Optional(),
		field.String("article_category_post_url").Optional(),
		field.JSON("article_content_selector", &custom.SourcePostSelector{}).Optional(),

		field.JSON("media_featured_post_selector", &custom.SourcePostSelector{}).Optional(),
		field.String("media_featured_post_url").Optional(),
		field.JSON("media_category_post_selector", &custom.SourcePostSelector{}).Optional(),
		field.String("media_category_post_url").Optional(),
		field.JSON("media_content_selector", &custom.SourcePostSelector{}).Optional(),

		field.JSON("article_categories", map[string]string{}).Optional(),
		field.JSON("media_categories", map[string]string{}).Optional(),
		field.String("description").NotEmpty(),
		field.String("language").Default("fr"),
		field.String("country").Default("ci"),
		field.Bool("status").Default(true),
		field.String("logo").NotEmpty(),
		field.String("name").Unique(),
		field.String("url"),
	}
}

// Edges of the Source.
func (Source) Edges() []ent.Edge {
	return nil
}
