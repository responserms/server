package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// Activation holds the schema definition for the Activation entity.
type Activation struct {
	ent.Schema
}

// Fields of the Activation.
func (Activation) Fields() []ent.Field {
	return []ent.Field{
		field.String("internal_comments").
			Optional().
			Nillable(),
		field.String("comments").
			Optional().
			Nillable(),
	}
}

// Mixin registers the mixins for the Activation entity.
func (Activation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Activation.
func (Activation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Annotations(entgql.Bind()),
		edge.To("actor", User.Type).
			Annotations(entgql.Bind()),
	}
}
