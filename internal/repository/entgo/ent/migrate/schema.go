// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// DialogsColumns holds the columns for the "dialogs" table.
	DialogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "user_dialog", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// DialogsTable holds the schema information for the "dialogs" table.
	DialogsTable = &schema.Table{
		Name:       "dialogs",
		Columns:    DialogsColumns,
		PrimaryKey: []*schema.Column{DialogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dialogs_users_dialog",
				Columns: []*schema.Column{DialogsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "attachment", Type: field.TypeString, Nullable: true},
		{Name: "system", Type: field.TypeString},
		{Name: "direction", Type: field.TypeString},
		{Name: "proactive", Type: field.TypeBool},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "dialog_message", Type: field.TypeInt, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "messages_dialogs_message",
				Columns: []*schema.Column{MessagesColumns[9]},

				RefColumns: []*schema.Column{DialogsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "user_projects", Type: field.TypeInt, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "projects_users_projects",
				Columns: []*schema.Column{ProjectsColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "service", Type: field.TypeBool},
		{Name: "teams_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "telegram_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "slack_id", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DialogsTable,
		MessagesTable,
		ProjectsTable,
		UsersTable,
	}
)

func init() {
	DialogsTable.ForeignKeys[0].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = DialogsTable
	ProjectsTable.ForeignKeys[0].RefTable = UsersTable
}
