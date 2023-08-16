package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("bug_id"),
		field.String("tag"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
