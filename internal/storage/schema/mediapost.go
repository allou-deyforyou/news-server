package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MediaPost holds the schema definition for the MediaPost entity.
type MediaPost struct {
	ent.Schema
}

// Fields of the MediaPostSchema.
func (MediaPost) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("status").Default(true),
		field.Bool("live").Default(false),
		field.Enum("type").Values("youtube", "audio", "video"),
		field.String("title").Optional(),
		field.String("logo"),
		field.String("source"),
		field.String("description"),
		field.String("image").Optional(),
		field.Time("date").Optional(),
		field.String("link").Optional(),
		field.String("content").Optional(),
	}
}

// Edges of the MediaPost.
func (MediaPost) Edges() []ent.Edge {
	return nil
}
