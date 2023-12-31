// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entedge/ent/bug"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Bug is the model entity for the Bug schema.
type Bug struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BugID holds the value of the "bug_id" field.
	BugID int `json:"bug_id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BugQuery when eager-loading is set.
	Edges        BugEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BugEdges holds the relations/edges for other nodes in the graph.
type BugEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e BugEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bug) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bug.FieldID, bug.FieldBugID:
			values[i] = new(sql.NullInt64)
		case bug.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bug fields.
func (b *Bug) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bug.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case bug.FieldBugID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bug_id", values[i])
			} else if value.Valid {
				b.BugID = int(value.Int64)
			}
		case bug.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bug.
// This includes values selected through modifiers, order, etc.
func (b *Bug) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryTags queries the "tags" edge of the Bug entity.
func (b *Bug) QueryTags() *TagQuery {
	return NewBugClient(b.config).QueryTags(b)
}

// Update returns a builder for updating this Bug.
// Note that you need to call Bug.Unwrap() before calling this method if this Bug
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bug) Update() *BugUpdateOne {
	return NewBugClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bug entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bug) Unwrap() *Bug {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bug is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bug) String() string {
	var builder strings.Builder
	builder.WriteString("Bug(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("bug_id=")
	builder.WriteString(fmt.Sprintf("%v", b.BugID))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(b.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Bugs is a parsable slice of Bug.
type Bugs []*Bug
