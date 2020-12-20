// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
	"github.com/facebook/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"password_hash,omitempty"`
	// Service holds the value of the "service" field.
	Service bool `json:"service,omitempty"`
	// TeamsID holds the value of the "teams_id" field.
	TeamsID *string `json:"teams_id,omitempty"`
	// TelegramID holds the value of the "telegram_id" field.
	TelegramID *string `json:"telegram_id,omitempty"`
	// SlackID holds the value of the "slack_id" field.
	SlackID *string `json:"slack_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Dialog holds the value of the dialog edge.
	Dialog *Dialog
	// Projects holds the value of the projects edge.
	Projects []*Project
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DialogOrErr returns the Dialog value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) DialogOrErr() (*Dialog, error) {
	if e.loadedTypes[0] {
		if e.Dialog == nil {
			// The edge dialog was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dialog.Label}
		}
		return e.Dialog, nil
	}
	return nil, &NotLoadedError{edge: "dialog"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // username
		&sql.NullString{}, // display_name
		&sql.NullString{}, // email
		&sql.NullString{}, // password_hash
		&sql.NullBool{},   // service
		&sql.NullString{}, // teams_id
		&sql.NullString{}, // telegram_id
		&sql.NullString{}, // slack_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		u.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		u.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[2])
	} else if value.Valid {
		u.Username = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field display_name", values[3])
	} else if value.Valid {
		u.DisplayName = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[4])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password_hash", values[5])
	} else if value.Valid {
		u.PasswordHash = value.String
	}
	if value, ok := values[6].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field service", values[6])
	} else if value.Valid {
		u.Service = value.Bool
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field teams_id", values[7])
	} else if value.Valid {
		u.TeamsID = new(string)
		*u.TeamsID = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field telegram_id", values[8])
	} else if value.Valid {
		u.TelegramID = new(string)
		*u.TelegramID = value.String
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field slack_id", values[9])
	} else if value.Valid {
		u.SlackID = new(string)
		*u.SlackID = value.String
	}
	return nil
}

// QueryDialog queries the dialog edge of the User.
func (u *User) QueryDialog() *DialogQuery {
	return (&UserClient{config: u.config}).QueryDialog(u)
}

// QueryProjects queries the projects edge of the User.
func (u *User) QueryProjects() *ProjectQuery {
	return (&UserClient{config: u.config}).QueryProjects(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password_hash=")
	builder.WriteString(u.PasswordHash)
	builder.WriteString(", service=")
	builder.WriteString(fmt.Sprintf("%v", u.Service))
	if v := u.TeamsID; v != nil {
		builder.WriteString(", teams_id=")
		builder.WriteString(*v)
	}
	if v := u.TelegramID; v != nil {
		builder.WriteString(", telegram_id=")
		builder.WriteString(*v)
	}
	if v := u.SlackID; v != nil {
		builder.WriteString(", slack_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
