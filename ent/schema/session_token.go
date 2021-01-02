package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// SessionToken holds the schema definition for the SessionToken entity.
type SessionToken struct {
	ent.Schema
}

// Fields of the SessionToken.
func (SessionToken) Fields() []ent.Field {
	return []ent.Field{
		field.Time("blocked_at").
			Optional().
			Nillable(),
		field.Time("expired_at").
			Optional().
			Nillable(),
	}
}

// Mixin provides shared configurations as a way to prevent code duplication.
func (SessionToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the SessionToken.
func (SessionToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("session_tokens").
			Unique().
			Annotations(entgql.Bind()),
	}
}
