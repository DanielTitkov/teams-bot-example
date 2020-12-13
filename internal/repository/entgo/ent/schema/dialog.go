package schema

import (
	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Dialog holds the schema definition for the Dialog entity.
type Dialog struct {
	ent.Schema
}

// Fields of the Dialog.
func (Dialog) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("meta", domain.DialogMeta{}).Optional(),
	}
}

// Edges of the Dialog.
func (Dialog) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("message", Message.Type),
		// belongs to
		edge.From("user", User.Type).Ref("dialog").Unique().Required(),
	}
}

func (Dialog) Indexes() []ent.Index {
	return nil
}

func (Dialog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
