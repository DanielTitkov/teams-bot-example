package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Optional(),
		field.String("attachment").Optional(),
		field.String("system"),
		field.String("direction"),
		field.Bool("proactive"),
		field.String("error").Optional().Nillable(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("dialog", Dialog.Type).Ref("message").Unique().Required(),
	}
}

func (Message) Indexes() []ent.Index {
	return nil
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
