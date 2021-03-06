// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/project"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the update_time field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetUsername sets the username field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetDisplayName sets the display_name field.
func (uc *UserCreate) SetDisplayName(s string) *UserCreate {
	uc.mutation.SetDisplayName(s)
	return uc
}

// SetNillableDisplayName sets the display_name field if the given value is not nil.
func (uc *UserCreate) SetNillableDisplayName(s *string) *UserCreate {
	if s != nil {
		uc.SetDisplayName(*s)
	}
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the email field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetPasswordHash sets the password_hash field.
func (uc *UserCreate) SetPasswordHash(s string) *UserCreate {
	uc.mutation.SetPasswordHash(s)
	return uc
}

// SetService sets the service field.
func (uc *UserCreate) SetService(b bool) *UserCreate {
	uc.mutation.SetService(b)
	return uc
}

// SetNillableService sets the service field if the given value is not nil.
func (uc *UserCreate) SetNillableService(b *bool) *UserCreate {
	if b != nil {
		uc.SetService(*b)
	}
	return uc
}

// SetTeamsID sets the teams_id field.
func (uc *UserCreate) SetTeamsID(s string) *UserCreate {
	uc.mutation.SetTeamsID(s)
	return uc
}

// SetNillableTeamsID sets the teams_id field if the given value is not nil.
func (uc *UserCreate) SetNillableTeamsID(s *string) *UserCreate {
	if s != nil {
		uc.SetTeamsID(*s)
	}
	return uc
}

// SetTelegramID sets the telegram_id field.
func (uc *UserCreate) SetTelegramID(s string) *UserCreate {
	uc.mutation.SetTelegramID(s)
	return uc
}

// SetNillableTelegramID sets the telegram_id field if the given value is not nil.
func (uc *UserCreate) SetNillableTelegramID(s *string) *UserCreate {
	if s != nil {
		uc.SetTelegramID(*s)
	}
	return uc
}

// SetSlackID sets the slack_id field.
func (uc *UserCreate) SetSlackID(s string) *UserCreate {
	uc.mutation.SetSlackID(s)
	return uc
}

// SetNillableSlackID sets the slack_id field if the given value is not nil.
func (uc *UserCreate) SetNillableSlackID(s *string) *UserCreate {
	if s != nil {
		uc.SetSlackID(*s)
	}
	return uc
}

// SetDialogID sets the dialog edge to Dialog by id.
func (uc *UserCreate) SetDialogID(id int) *UserCreate {
	uc.mutation.SetDialogID(id)
	return uc
}

// SetNillableDialogID sets the dialog edge to Dialog by id if the given value is not nil.
func (uc *UserCreate) SetNillableDialogID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetDialogID(*id)
	}
	return uc
}

// SetDialog sets the dialog edge to Dialog.
func (uc *UserCreate) SetDialog(d *Dialog) *UserCreate {
	return uc.SetDialogID(d.ID)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (uc *UserCreate) AddProjectIDs(ids ...int) *UserCreate {
	uc.mutation.AddProjectIDs(ids...)
	return uc
}

// AddProjects adds the projects edges to Project.
func (uc *UserCreate) AddProjects(p ...*Project) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddProjectIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) preSave() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("ent: missing required field \"username\"")}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if _, ok := uc.mutation.PasswordHash(); !ok {
		return &ValidationError{Name: "password_hash", err: errors.New("ent: missing required field \"password_hash\"")}
	}
	if _, ok := uc.mutation.Service(); !ok {
		v := user.DefaultService
		uc.mutation.SetService(v)
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateTime,
		})
		u.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
		u.UpdateTime = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
		u.Username = value
	}
	if value, ok := uc.mutation.DisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
		u.DisplayName = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		u.Email = &value
	}
	if value, ok := uc.mutation.PasswordHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
		u.PasswordHash = value
	}
	if value, ok := uc.mutation.Service(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldService,
		})
		u.Service = value
	}
	if value, ok := uc.mutation.TeamsID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTeamsID,
		})
		u.TeamsID = &value
	}
	if value, ok := uc.mutation.TelegramID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTelegramID,
		})
		u.TelegramID = &value
	}
	if value, ok := uc.mutation.SlackID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSlackID,
		})
		u.SlackID = &value
	}
	if nodes := uc.mutation.DialogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.DialogTable,
			Columns: []string{user.DialogColumn},
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
	if nodes := uc.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}

// UserCreateBulk is the builder for creating a bulk of User entities.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
