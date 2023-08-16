// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entedge/ent/bug"
	"entedge/ent/predicate"
	"entedge/ent/tag"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BugUpdate is the builder for updating Bug entities.
type BugUpdate struct {
	config
	hooks    []Hook
	mutation *BugMutation
}

// Where appends a list predicates to the BugUpdate builder.
func (bu *BugUpdate) Where(ps ...predicate.Bug) *BugUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBugID sets the "bug_id" field.
func (bu *BugUpdate) SetBugID(i int) *BugUpdate {
	bu.mutation.ResetBugID()
	bu.mutation.SetBugID(i)
	return bu
}

// AddBugID adds i to the "bug_id" field.
func (bu *BugUpdate) AddBugID(i int) *BugUpdate {
	bu.mutation.AddBugID(i)
	return bu
}

// SetDescription sets the "description" field.
func (bu *BugUpdate) SetDescription(s string) *BugUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (bu *BugUpdate) AddTagIDs(ids ...int) *BugUpdate {
	bu.mutation.AddTagIDs(ids...)
	return bu
}

// AddTags adds the "tags" edges to the Tag entity.
func (bu *BugUpdate) AddTags(t ...*Tag) *BugUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTagIDs(ids...)
}

// Mutation returns the BugMutation object of the builder.
func (bu *BugUpdate) Mutation() *BugMutation {
	return bu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (bu *BugUpdate) ClearTags() *BugUpdate {
	bu.mutation.ClearTags()
	return bu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (bu *BugUpdate) RemoveTagIDs(ids ...int) *BugUpdate {
	bu.mutation.RemoveTagIDs(ids...)
	return bu
}

// RemoveTags removes "tags" edges to Tag entities.
func (bu *BugUpdate) RemoveTags(t ...*Tag) *BugUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BugUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BugUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BugUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BugUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BugUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bug.Table, bug.Columns, sqlgraph.NewFieldSpec(bug.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BugID(); ok {
		_spec.SetField(bug.FieldBugID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedBugID(); ok {
		_spec.AddField(bug.FieldBugID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(bug.FieldDescription, field.TypeString, value)
	}
	if bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bug.TagsTable,
			Columns: []string{bug.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bug.TagsTable,
			Columns: []string{bug.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bug.TagsTable,
			Columns: []string{bug.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bug.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BugUpdateOne is the builder for updating a single Bug entity.
type BugUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BugMutation
}

// SetBugID sets the "bug_id" field.
func (buo *BugUpdateOne) SetBugID(i int) *BugUpdateOne {
	buo.mutation.ResetBugID()
	buo.mutation.SetBugID(i)
	return buo
}

// AddBugID adds i to the "bug_id" field.
func (buo *BugUpdateOne) AddBugID(i int) *BugUpdateOne {
	buo.mutation.AddBugID(i)
	return buo
}

// SetDescription sets the "description" field.
func (buo *BugUpdateOne) SetDescription(s string) *BugUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (buo *BugUpdateOne) AddTagIDs(ids ...int) *BugUpdateOne {
	buo.mutation.AddTagIDs(ids...)
	return buo
}

// AddTags adds the "tags" edges to the Tag entity.
func (buo *BugUpdateOne) AddTags(t ...*Tag) *BugUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTagIDs(ids...)
}

// Mutation returns the BugMutation object of the builder.
func (buo *BugUpdateOne) Mutation() *BugMutation {
	return buo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (buo *BugUpdateOne) ClearTags() *BugUpdateOne {
	buo.mutation.ClearTags()
	return buo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (buo *BugUpdateOne) RemoveTagIDs(ids ...int) *BugUpdateOne {
	buo.mutation.RemoveTagIDs(ids...)
	return buo
}

// RemoveTags removes "tags" edges to Tag entities.
func (buo *BugUpdateOne) RemoveTags(t ...*Tag) *BugUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTagIDs(ids...)
}

// Where appends a list predicates to the BugUpdate builder.
func (buo *BugUpdateOne) Where(ps ...predicate.Bug) *BugUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BugUpdateOne) Select(field string, fields ...string) *BugUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bug entity.
func (buo *BugUpdateOne) Save(ctx context.Context) (*Bug, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BugUpdateOne) SaveX(ctx context.Context) *Bug {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BugUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BugUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BugUpdateOne) sqlSave(ctx context.Context) (_node *Bug, err error) {
	_spec := sqlgraph.NewUpdateSpec(bug.Table, bug.Columns, sqlgraph.NewFieldSpec(bug.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bug.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bug.FieldID)
		for _, f := range fields {
			if !bug.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bug.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.BugID(); ok {
		_spec.SetField(bug.FieldBugID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedBugID(); ok {
		_spec.AddField(bug.FieldBugID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(bug.FieldDescription, field.TypeString, value)
	}
	if buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bug.TagsTable,
			Columns: []string{bug.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bug.TagsTable,
			Columns: []string{bug.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bug.TagsTable,
			Columns: []string{bug.TagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bug{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bug.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
