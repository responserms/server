package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip_address"),
		field.String("browser_name"),
		field.String("browser_version"),
		field.String("device_os"),
		field.String("device_type"),
		field.Time("terminated_at").
			Optional().
			Nillable(),
	}
}

// Mixin for session.
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("token", Token.Type).
			Ref("session").
			Unique().
			Annotations(entgql.Bind()),
		edge.From("user", User.Type).
			Ref("sessions").
			Annotations(entgql.Bind()).
			Unique(),
	}
}

// Indexes of the Sessopn.
func (Session) Indexes() []ent.Index {
	return []ent.Index{}
}
