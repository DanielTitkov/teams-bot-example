// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/message"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// DialogUpdate is the builder for updating Dialog entities.
type DialogUpdate struct {
	config
	hooks      []Hook
	mutation   *DialogMutation
	predicates []predicate.Dialog
}

// Where adds a new predicate for the builder.
func (du *DialogUpdate) Where(ps ...predicate.Dialog) *DialogUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetMeta sets the meta field.
func (du *DialogUpdate) SetMeta(dm domain.DialogMeta) *DialogUpdate {
	du.mutation.SetMeta(dm)
	return du
}

// SetNillableMeta sets the meta field if the given value is not nil.
func (du *DialogUpdate) SetNillableMeta(dm *domain.DialogMeta) *DialogUpdate {
	if dm != nil {
		du.SetMeta(*dm)
	}
	return du
}

// ClearMeta clears the value of meta.
func (du *DialogUpdate) ClearMeta() *DialogUpdate {
	du.mutation.ClearMeta()
	return du
}

// AddMessageIDs adds the message edge to Message by ids.
func (du *DialogUpdate) AddMessageIDs(ids ...int) *DialogUpdate {
	du.mutation.AddMessageIDs(ids...)
	return du
}

// AddMessage adds the message edges to Message.
func (du *DialogUpdate) AddMessage(m ...*Message) *DialogUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return du.AddMessageIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (du *DialogUpdate) SetUserID(id int) *DialogUpdate {
	du.mutation.SetUserID(id)
	return du
}

// SetUser sets the user edge to User.
func (du *DialogUpdate) SetUser(u *User) *DialogUpdate {
	return du.SetUserID(u.ID)
}

// Mutation returns the DialogMutation object of the builder.
func (du *DialogUpdate) Mutation() *DialogMutation {
	return du.mutation
}

// RemoveMessageIDs removes the message edge to Message by ids.
func (du *DialogUpdate) RemoveMessageIDs(ids ...int) *DialogUpdate {
	du.mutation.RemoveMessageIDs(ids...)
	return du
}

// RemoveMessage removes message edges to Message.
func (du *DialogUpdate) RemoveMessage(m ...*Message) *DialogUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return du.RemoveMessageIDs(ids...)
}

// ClearUser clears the user edge to User.
func (du *DialogUpdate) ClearUser() *DialogUpdate {
	du.mutation.ClearUser()
	return du
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DialogUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := du.mutation.UpdateTime(); !ok {
		v := dialog.UpdateDefaultUpdateTime()
		du.mutation.SetUpdateTime(v)
	}

	if _, ok := du.mutation.UserID(); du.mutation.UserCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DialogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DialogUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DialogUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DialogUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DialogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dialog.Table,
			Columns: dialog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dialog.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dialog.FieldUpdateTime,
		})
	}
	if value, ok := du.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dialog.FieldMeta,
		})
	}
	if du.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dialog.FieldMeta,
		})
	}
	if nodes := du.mutation.RemovedMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dialog.MessageTable,
			Columns: []string{dialog.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dialog.MessageTable,
			Columns: []string{dialog.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dialog.UserTable,
			Columns: []string{dialog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dialog.UserTable,
			Columns: []string{dialog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dialog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DialogUpdateOne is the builder for updating a single Dialog entity.
type DialogUpdateOne struct {
	config
	hooks    []Hook
	mutation *DialogMutation
}

// SetMeta sets the meta field.
func (duo *DialogUpdateOne) SetMeta(dm domain.DialogMeta) *DialogUpdateOne {
	duo.mutation.SetMeta(dm)
	return duo
}

// SetNillableMeta sets the meta field if the given value is not nil.
func (duo *DialogUpdateOne) SetNillableMeta(dm *domain.DialogMeta) *DialogUpdateOne {
	if dm != nil {
		duo.SetMeta(*dm)
	}
	return duo
}

// ClearMeta clears the value of meta.
func (duo *DialogUpdateOne) ClearMeta() *DialogUpdateOne {
	duo.mutation.ClearMeta()
	return duo
}

// AddMessageIDs adds the message edge to Message by ids.
func (duo *DialogUpdateOne) AddMessageIDs(ids ...int) *DialogUpdateOne {
	duo.mutation.AddMessageIDs(ids...)
	return duo
}

// AddMessage adds the message edges to Message.
func (duo *DialogUpdateOne) AddMessage(m ...*Message) *DialogUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duo.AddMessageIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (duo *DialogUpdateOne) SetUserID(id int) *DialogUpdateOne {
	duo.mutation.SetUserID(id)
	return duo
}

// SetUser sets the user edge to User.
func (duo *DialogUpdateOne) SetUser(u *User) *DialogUpdateOne {
	return duo.SetUserID(u.ID)
}

// Mutation returns the DialogMutation object of the builder.
func (duo *DialogUpdateOne) Mutation() *DialogMutation {
	return duo.mutation
}

// RemoveMessageIDs removes the message edge to Message by ids.
func (duo *DialogUpdateOne) RemoveMessageIDs(ids ...int) *DialogUpdateOne {
	duo.mutation.RemoveMessageIDs(ids...)
	return duo
}

// RemoveMessage removes message edges to Message.
func (duo *DialogUpdateOne) RemoveMessage(m ...*Message) *DialogUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duo.RemoveMessageIDs(ids...)
}

// ClearUser clears the user edge to User.
func (duo *DialogUpdateOne) ClearUser() *DialogUpdateOne {
	duo.mutation.ClearUser()
	return duo
}

// Save executes the query and returns the updated entity.
func (duo *DialogUpdateOne) Save(ctx context.Context) (*Dialog, error) {
	if _, ok := duo.mutation.UpdateTime(); !ok {
		v := dialog.UpdateDefaultUpdateTime()
		duo.mutation.SetUpdateTime(v)
	}

	if _, ok := duo.mutation.UserID(); duo.mutation.UserCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"user\"")
	}
	var (
		err  error
		node *Dialog
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DialogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DialogUpdateOne) SaveX(ctx context.Context) *Dialog {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DialogUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DialogUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DialogUpdateOne) sqlSave(ctx context.Context) (d *Dialog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dialog.Table,
			Columns: dialog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dialog.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dialog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dialog.FieldUpdateTime,
		})
	}
	if value, ok := duo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dialog.FieldMeta,
		})
	}
	if duo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dialog.FieldMeta,
		})
	}
	if nodes := duo.mutation.RemovedMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dialog.MessageTable,
			Columns: []string{dialog.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dialog.MessageTable,
			Columns: []string{dialog.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dialog.UserTable,
			Columns: []string{dialog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   dialog.UserTable,
			Columns: []string{dialog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Dialog{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dialog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}