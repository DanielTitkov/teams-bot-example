// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/message"
	"github.com/facebook/ent/dialect/sql"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Attachment holds the value of the "attachment" field.
	Attachment string `json:"attachment,omitempty"`
	// PayloadType holds the value of the "payload_type" field.
	PayloadType string `json:"payload_type,omitempty"`
	// PayloadValue holds the value of the "payload_value" field.
	PayloadValue string `json:"payload_value,omitempty"`
	// System holds the value of the "system" field.
	System string `json:"system,omitempty"`
	// Direction holds the value of the "direction" field.
	Direction string `json:"direction,omitempty"`
	// Proactive holds the value of the "proactive" field.
	Proactive bool `json:"proactive,omitempty"`
	// Error holds the value of the "error" field.
	Error *string `json:"error,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges          MessageEdges `json:"edges"`
	dialog_message *int
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// Dialog holds the value of the dialog edge.
	Dialog *Dialog
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DialogOrErr returns the Dialog value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) DialogOrErr() (*Dialog, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // text
		&sql.NullString{}, // attachment
		&sql.NullString{}, // payload_type
		&sql.NullString{}, // payload_value
		&sql.NullString{}, // system
		&sql.NullString{}, // direction
		&sql.NullBool{},   // proactive
		&sql.NullString{}, // error
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Message) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // dialog_message
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(values ...interface{}) error {
	if m, n := len(values), len(message.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		m.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		m.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field text", values[2])
	} else if value.Valid {
		m.Text = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field attachment", values[3])
	} else if value.Valid {
		m.Attachment = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field payload_type", values[4])
	} else if value.Valid {
		m.PayloadType = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field payload_value", values[5])
	} else if value.Valid {
		m.PayloadValue = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field system", values[6])
	} else if value.Valid {
		m.System = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field direction", values[7])
	} else if value.Valid {
		m.Direction = value.String
	}
	if value, ok := values[8].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field proactive", values[8])
	} else if value.Valid {
		m.Proactive = value.Bool
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field error", values[9])
	} else if value.Valid {
		m.Error = new(string)
		*m.Error = value.String
	}
	values = values[10:]
	if len(values) == len(message.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field dialog_message", value)
		} else if value.Valid {
			m.dialog_message = new(int)
			*m.dialog_message = int(value.Int64)
		}
	}
	return nil
}

// QueryDialog queries the dialog edge of the Message.
func (m *Message) QueryDialog() *DialogQuery {
	return (&MessageClient{config: m.config}).QueryDialog(m)
}

// Update returns a builder for updating this Message.
// Note that, you need to call Message.Unwrap() before calling this method, if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return (&MessageClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", text=")
	builder.WriteString(m.Text)
	builder.WriteString(", attachment=")
	builder.WriteString(m.Attachment)
	builder.WriteString(", payload_type=")
	builder.WriteString(m.PayloadType)
	builder.WriteString(", payload_value=")
	builder.WriteString(m.PayloadValue)
	builder.WriteString(", system=")
	builder.WriteString(m.System)
	builder.WriteString(", direction=")
	builder.WriteString(m.Direction)
	builder.WriteString(", proactive=")
	builder.WriteString(fmt.Sprintf("%v", m.Proactive))
	if v := m.Error; v != nil {
		builder.WriteString(", error=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
