// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/dialog"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/message"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/project"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/schema"
	"github.com/DanielTitkov/teams-bot-example/internal/repository/entgo/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	dialogMixin := schema.Dialog{}.Mixin()
	dialogMixinFields0 := dialogMixin[0].Fields()
	dialogFields := schema.Dialog{}.Fields()
	_ = dialogFields
	// dialogDescCreateTime is the schema descriptor for create_time field.
	dialogDescCreateTime := dialogMixinFields0[0].Descriptor()
	// dialog.DefaultCreateTime holds the default value on creation for the create_time field.
	dialog.DefaultCreateTime = dialogDescCreateTime.Default.(func() time.Time)
	// dialogDescUpdateTime is the schema descriptor for update_time field.
	dialogDescUpdateTime := dialogMixinFields0[1].Descriptor()
	// dialog.DefaultUpdateTime holds the default value on creation for the update_time field.
	dialog.DefaultUpdateTime = dialogDescUpdateTime.Default.(func() time.Time)
	// dialog.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	dialog.UpdateDefaultUpdateTime = dialogDescUpdateTime.UpdateDefault.(func() time.Time)
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreateTime is the schema descriptor for create_time field.
	messageDescCreateTime := messageMixinFields0[0].Descriptor()
	// message.DefaultCreateTime holds the default value on creation for the create_time field.
	message.DefaultCreateTime = messageDescCreateTime.Default.(func() time.Time)
	// messageDescUpdateTime is the schema descriptor for update_time field.
	messageDescUpdateTime := messageMixinFields0[1].Descriptor()
	// message.DefaultUpdateTime holds the default value on creation for the update_time field.
	message.DefaultUpdateTime = messageDescUpdateTime.Default.(func() time.Time)
	// message.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	message.UpdateDefaultUpdateTime = messageDescUpdateTime.UpdateDefault.(func() time.Time)
	projectMixin := schema.Project{}.Mixin()
	projectMixinFields0 := projectMixin[0].Fields()
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreateTime is the schema descriptor for create_time field.
	projectDescCreateTime := projectMixinFields0[0].Descriptor()
	// project.DefaultCreateTime holds the default value on creation for the create_time field.
	project.DefaultCreateTime = projectDescCreateTime.Default.(func() time.Time)
	// projectDescUpdateTime is the schema descriptor for update_time field.
	projectDescUpdateTime := projectMixinFields0[1].Descriptor()
	// project.DefaultUpdateTime holds the default value on creation for the update_time field.
	project.DefaultUpdateTime = projectDescUpdateTime.Default.(func() time.Time)
	// project.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	project.UpdateDefaultUpdateTime = projectDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescService is the schema descriptor for service field.
	userDescService := userFields[4].Descriptor()
	// user.DefaultService holds the default value on creation for the service field.
	user.DefaultService = userDescService.Default.(bool)
}
