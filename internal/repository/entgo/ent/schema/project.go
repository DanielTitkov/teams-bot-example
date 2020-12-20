package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Time("due_date"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("user", User.Type).Ref("projects").Unique().Required(),
	}
}

func (Project) Indexes() []ent.Index {
	return nil
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
