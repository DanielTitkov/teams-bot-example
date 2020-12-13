// Code generated by entc, DO NOT EDIT.

package dialog

import (
	"time"
)

const (
	// Label holds the string label denoting the dialog type in the database.
	Label = "dialog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"

	// EdgeMessage holds the string denoting the message edge name in mutations.
	EdgeMessage = "message"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the dialog in the database.
	Table = "dialogs"
	// MessageTable is the table the holds the message relation/edge.
	MessageTable = "messages"
	// MessageInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessageInverseTable = "messages"
	// MessageColumn is the table column denoting the message relation/edge.
	MessageColumn = "dialog_message"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "dialogs"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_dialog"
)

// Columns holds all SQL columns for dialog fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldMeta,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Dialog type.
var ForeignKeys = []string{
	"user_dialog",
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
)
