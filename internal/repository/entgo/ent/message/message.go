// Code generated by entc, DO NOT EDIT.

package message

import (
	"time"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldAttachment holds the string denoting the attachment field in the database.
	FieldAttachment = "attachment"
	// FieldPayloadType holds the string denoting the payload_type field in the database.
	FieldPayloadType = "payload_type"
	// FieldPayloadValue holds the string denoting the payload_value field in the database.
	FieldPayloadValue = "payload_value"
	// FieldSystem holds the string denoting the system field in the database.
	FieldSystem = "system"
	// FieldDirection holds the string denoting the direction field in the database.
	FieldDirection = "direction"
	// FieldProactive holds the string denoting the proactive field in the database.
	FieldProactive = "proactive"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"

	// EdgeDialog holds the string denoting the dialog edge name in mutations.
	EdgeDialog = "dialog"

	// Table holds the table name of the message in the database.
	Table = "messages"
	// DialogTable is the table the holds the dialog relation/edge.
	DialogTable = "messages"
	// DialogInverseTable is the table name for the Dialog entity.
	// It exists in this package in order to avoid circular dependency with the "dialog" package.
	DialogInverseTable = "dialogs"
	// DialogColumn is the table column denoting the dialog relation/edge.
	DialogColumn = "dialog_message"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldText,
	FieldAttachment,
	FieldPayloadType,
	FieldPayloadValue,
	FieldSystem,
	FieldDirection,
	FieldProactive,
	FieldError,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Message type.
var ForeignKeys = []string{
	"dialog_message",
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
)
