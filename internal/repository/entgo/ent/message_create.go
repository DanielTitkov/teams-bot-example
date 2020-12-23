// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/message"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (mc *MessageCreate) SetCreateTime(t time.Time) *MessageCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreateTime(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the update_time field.
func (mc *MessageCreate) SetUpdateTime(t time.Time) *MessageCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdateTime(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetText sets the text field.
func (mc *MessageCreate) SetText(s string) *MessageCreate {
	mc.mutation.SetText(s)
	return mc
}

// SetNillableText sets the text field if the given value is not nil.
func (mc *MessageCreate) SetNillableText(s *string) *MessageCreate {
	if s != nil {
		mc.SetText(*s)
	}
	return mc
}

// SetAttachment sets the attachment field.
func (mc *MessageCreate) SetAttachment(s string) *MessageCreate {
	mc.mutation.SetAttachment(s)
	return mc
}

// SetNillableAttachment sets the attachment field if the given value is not nil.
func (mc *MessageCreate) SetNillableAttachment(s *string) *MessageCreate {
	if s != nil {
		mc.SetAttachment(*s)
	}
	return mc
}

// SetSystem sets the system field.
func (mc *MessageCreate) SetSystem(s string) *MessageCreate {
	mc.mutation.SetSystem(s)
	return mc
}

// SetDirection sets the direction field.
func (mc *MessageCreate) SetDirection(s string) *MessageCreate {
	mc.mutation.SetDirection(s)
	return mc
}

// SetProactive sets the proactive field.
func (mc *MessageCreate) SetProactive(b bool) *MessageCreate {
	mc.mutation.SetProactive(b)
	return mc
}

// SetError sets the error field.
func (mc *MessageCreate) SetError(s string) *MessageCreate {
	mc.mutation.SetError(s)
	return mc
}

// SetNillableError sets the error field if the given value is not nil.
func (mc *MessageCreate) SetNillableError(s *string) *MessageCreate {
	if s != nil {
		mc.SetError(*s)
	}
	return mc
}

// SetDialogID sets the dialog edge to Dialog by id.
func (mc *MessageCreate) SetDialogID(id int) *MessageCreate {
	mc.mutation.SetDialogID(id)
	return mc
}

// SetDialog sets the dialog edge to Dialog.
func (mc *MessageCreate) SetDialog(d *Dialog) *MessageCreate {
	return mc.SetDialogID(d.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	if err := mc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Message
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MessageCreate) preSave() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := message.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := message.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
	if _, ok := mc.mutation.System(); !ok {
		return &ValidationError{Name: "system", err: errors.New("ent: missing required field \"system\"")}
	}
	if _, ok := mc.mutation.Direction(); !ok {
		return &ValidationError{Name: "direction", err: errors.New("ent: missing required field \"direction\"")}
	}
	if _, ok := mc.mutation.Proactive(); !ok {
		return &ValidationError{Name: "proactive", err: errors.New("ent: missing required field \"proactive\"")}
	}
	if _, ok := mc.mutation.DialogID(); !ok {
		return &ValidationError{Name: "dialog", err: errors.New("ent: missing required edge \"dialog\"")}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		m     = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldCreateTime,
		})
		m.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldUpdateTime,
		})
		m.UpdateTime = value
	}
	if value, ok := mc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldText,
		})
		m.Text = value
	}
	if value, ok := mc.mutation.Attachment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldAttachment,
		})
		m.Attachment = value
	}
	if value, ok := mc.mutation.System(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldSystem,
		})
		m.System = value
	}
	if value, ok := mc.mutation.Direction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldDirection,
		})
		m.Direction = value
	}
	if value, ok := mc.mutation.Proactive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldProactive,
		})
		m.Proactive = value
	}
	if value, ok := mc.mutation.Error(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldError,
		})
		m.Error = &value
	}
	if nodes := mc.mutation.DialogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.DialogTable,
			Columns: []string{message.DialogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dialog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}

// MessageCreateBulk is the builder for creating a bulk of Message entities.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*MessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
