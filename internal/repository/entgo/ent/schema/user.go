package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty(),
		field.String("display_name").Optional(),
		field.String("email").Unique().Optional().Nillable(),
		field.String("password_hash"),
		field.Bool("service").Default(false),
		field.String("teams_id").Optional().Nillable().Unique(),
		field.String("telegram_id").Optional().Nillable().Unique(),
		field.String("slack_id").Optional().Nillable().Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dialog", Dialog.Type).Unique(),
		edge.To("projects", Project.Type),
	}
}

func (User) Indexes() []ent.Index {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
