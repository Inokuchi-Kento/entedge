package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bug holds the schema definition for the Bug entity.
type Bug struct {
	ent.Schema
}

// Fields of the Bug.
func (Bug) Fields() []ent.Field {
	return []ent.Field{
		field.Int("bug_id").Unique(),
		field.String("description"),
	}
}

// Edges of the Bug.
func (Bug) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).
			StorageKey(edge.Column("bug_id")),
	}
}
