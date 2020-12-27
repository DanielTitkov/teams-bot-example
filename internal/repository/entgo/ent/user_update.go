// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/project"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetUsername sets the username field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetDisplayName sets the display_name field.
func (uu *UserUpdate) SetDisplayName(s string) *UserUpdate {
	uu.mutation.SetDisplayName(s)
	return uu
}

// SetNillableDisplayName sets the display_name field if the given value is not nil.
func (uu *UserUpdate) SetNillableDisplayName(s *string) *UserUpdate {
	if s != nil {
		uu.SetDisplayName(*s)
	}
	return uu
}

// ClearDisplayName clears the value of display_name.
func (uu *UserUpdate) ClearDisplayName() *UserUpdate {
	uu.mutation.ClearDisplayName()
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the email field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of email.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetPasswordHash sets the password_hash field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.mutation.SetPasswordHash(s)
	return uu
}

// SetService sets the service field.
func (uu *UserUpdate) SetService(b bool) *UserUpdate {
	uu.mutation.SetService(b)
	return uu
}

// SetNillableService sets the service field if the given value is not nil.
func (uu *UserUpdate) SetNillableService(b *bool) *UserUpdate {
	if b != nil {
		uu.SetService(*b)
	}
	return uu
}

// SetTeamsID sets the teams_id field.
func (uu *UserUpdate) SetTeamsID(s string) *UserUpdate {
	uu.mutation.SetTeamsID(s)
	return uu
}

// SetNillableTeamsID sets the teams_id field if the given value is not nil.
func (uu *UserUpdate) SetNillableTeamsID(s *string) *UserUpdate {
	if s != nil {
		uu.SetTeamsID(*s)
	}
	return uu
}

// ClearTeamsID clears the value of teams_id.
func (uu *UserUpdate) ClearTeamsID() *UserUpdate {
	uu.mutation.ClearTeamsID()
	return uu
}

// SetTelegramID sets the telegram_id field.
func (uu *UserUpdate) SetTelegramID(s string) *UserUpdate {
	uu.mutation.SetTelegramID(s)
	return uu
}

// SetNillableTelegramID sets the telegram_id field if the given value is not nil.
func (uu *UserUpdate) SetNillableTelegramID(s *string) *UserUpdate {
	if s != nil {
		uu.SetTelegramID(*s)
	}
	return uu
}

// ClearTelegramID clears the value of telegram_id.
func (uu *UserUpdate) ClearTelegramID() *UserUpdate {
	uu.mutation.ClearTelegramID()
	return uu
}

// SetSlackID sets the slack_id field.
func (uu *UserUpdate) SetSlackID(s string) *UserUpdate {
	uu.mutation.SetSlackID(s)
	return uu
}

// SetNillableSlackID sets the slack_id field if the given value is not nil.
func (uu *UserUpdate) SetNillableSlackID(s *string) *UserUpdate {
	if s != nil {
		uu.SetSlackID(*s)
	}
	return uu
}

// ClearSlackID clears the value of slack_id.
func (uu *UserUpdate) ClearSlackID() *UserUpdate {
	uu.mutation.ClearSlackID()
	return uu
}

// SetDialogID sets the dialog edge to Dialog by id.
func (uu *UserUpdate) SetDialogID(id int) *UserUpdate {
	uu.mutation.SetDialogID(id)
	return uu
}

// SetNillableDialogID sets the dialog edge to Dialog by id if the given value is not nil.
func (uu *UserUpdate) SetNillableDialogID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetDialogID(*id)
	}
	return uu
}

// SetDialog sets the dialog edge to Dialog.
func (uu *UserUpdate) SetDialog(d *Dialog) *UserUpdate {
	return uu.SetDialogID(d.ID)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (uu *UserUpdate) AddProjectIDs(ids ...int) *UserUpdate {
	uu.mutation.AddProjectIDs(ids...)
	return uu
}

// AddProjects adds the projects edges to Project.
func (uu *UserUpdate) AddProjects(p ...*Project) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddProjectIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearDialog clears the dialog edge to Dialog.
func (uu *UserUpdate) ClearDialog() *UserUpdate {
	uu.mutation.ClearDialog()
	return uu
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (uu *UserUpdate) RemoveProjectIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveProjectIDs(ids...)
	return uu
}

// RemoveProjects removes projects edges to Project.
func (uu *UserUpdate) RemoveProjects(p ...*Project) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return 0, &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
	}
	if uu.mutation.DisplayNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldDisplayName,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if uu.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
	}
	if value, ok := uu.mutation.Service(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldService,
		})
	}
	if value, ok := uu.mutation.TeamsID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTeamsID,
		})
	}
	if uu.mutation.TeamsIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTeamsID,
		})
	}
	if value, ok := uu.mutation.TelegramID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTelegramID,
		})
	}
	if uu.mutation.TelegramIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTelegramID,
		})
	}
	if value, ok := uu.mutation.SlackID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSlackID,
		})
	}
	if uu.mutation.SlackIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSlackID,
		})
	}
	if uu.mutation.DialogCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DialogIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedProjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the username field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetDisplayName sets the display_name field.
func (uuo *UserUpdateOne) SetDisplayName(s string) *UserUpdateOne {
	uuo.mutation.SetDisplayName(s)
	return uuo
}

// SetNillableDisplayName sets the display_name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDisplayName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDisplayName(*s)
	}
	return uuo
}

// ClearDisplayName clears the value of display_name.
func (uuo *UserUpdateOne) ClearDisplayName() *UserUpdateOne {
	uuo.mutation.ClearDisplayName()
	return uuo
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the email field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of email.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetPasswordHash sets the password_hash field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(s)
	return uuo
}

// SetService sets the service field.
func (uuo *UserUpdateOne) SetService(b bool) *UserUpdateOne {
	uuo.mutation.SetService(b)
	return uuo
}

// SetNillableService sets the service field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableService(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetService(*b)
	}
	return uuo
}

// SetTeamsID sets the teams_id field.
func (uuo *UserUpdateOne) SetTeamsID(s string) *UserUpdateOne {
	uuo.mutation.SetTeamsID(s)
	return uuo
}

// SetNillableTeamsID sets the teams_id field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTeamsID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTeamsID(*s)
	}
	return uuo
}

// ClearTeamsID clears the value of teams_id.
func (uuo *UserUpdateOne) ClearTeamsID() *UserUpdateOne {
	uuo.mutation.ClearTeamsID()
	return uuo
}

// SetTelegramID sets the telegram_id field.
func (uuo *UserUpdateOne) SetTelegramID(s string) *UserUpdateOne {
	uuo.mutation.SetTelegramID(s)
	return uuo
}

// SetNillableTelegramID sets the telegram_id field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTelegramID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTelegramID(*s)
	}
	return uuo
}

// ClearTelegramID clears the value of telegram_id.
func (uuo *UserUpdateOne) ClearTelegramID() *UserUpdateOne {
	uuo.mutation.ClearTelegramID()
	return uuo
}

// SetSlackID sets the slack_id field.
func (uuo *UserUpdateOne) SetSlackID(s string) *UserUpdateOne {
	uuo.mutation.SetSlackID(s)
	return uuo
}

// SetNillableSlackID sets the slack_id field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSlackID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSlackID(*s)
	}
	return uuo
}

// ClearSlackID clears the value of slack_id.
func (uuo *UserUpdateOne) ClearSlackID() *UserUpdateOne {
	uuo.mutation.ClearSlackID()
	return uuo
}

// SetDialogID sets the dialog edge to Dialog by id.
func (uuo *UserUpdateOne) SetDialogID(id int) *UserUpdateOne {
	uuo.mutation.SetDialogID(id)
	return uuo
}

// SetNillableDialogID sets the dialog edge to Dialog by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDialogID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetDialogID(*id)
	}
	return uuo
}

// SetDialog sets the dialog edge to Dialog.
func (uuo *UserUpdateOne) SetDialog(d *Dialog) *UserUpdateOne {
	return uuo.SetDialogID(d.ID)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (uuo *UserUpdateOne) AddProjectIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddProjectIDs(ids...)
	return uuo
}

// AddProjects adds the projects edges to Project.
func (uuo *UserUpdateOne) AddProjects(p ...*Project) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddProjectIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearDialog clears the dialog edge to Dialog.
func (uuo *UserUpdateOne) ClearDialog() *UserUpdateOne {
	uuo.mutation.ClearDialog()
	return uuo
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (uuo *UserUpdateOne) RemoveProjectIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveProjectIDs(ids...)
	return uuo
}

// RemoveProjects removes projects edges to Project.
func (uuo *UserUpdateOne) RemoveProjects(p ...*Project) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return nil, &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}

	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uuo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
	}
	if uuo.mutation.DisplayNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldDisplayName,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if uuo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
	}
	if value, ok := uuo.mutation.Service(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldService,
		})
	}
	if value, ok := uuo.mutation.TeamsID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTeamsID,
		})
	}
	if uuo.mutation.TeamsIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTeamsID,
		})
	}
	if value, ok := uuo.mutation.TelegramID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTelegramID,
		})
	}
	if uuo.mutation.TelegramIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTelegramID,
		})
	}
	if value, ok := uuo.mutation.SlackID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSlackID,
		})
	}
	if uuo.mutation.SlackIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldSlackID,
		})
	}
	if uuo.mutation.DialogCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DialogIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
