package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		// bytea型は、PostgreSQLのみでサポートされています。
		field.Bytes("data"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
