package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("total_minutes").
			Default(0).
			Annotations(entgql.OrderField("TOTAL_MINUTES")),
		field.Time("session_started_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("SESSION_STARTED_AT")),
		field.Time("session_ended_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("SESSION_ENDED_AT")),
		field.Time("last_seen_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("LAST_SEEN_AT")),
	}
}

func (Player) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		MetadataMixin{},
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servers", Server.Type).
			Annotations(entgql.Bind()),
		edge.To("user", User.Type).
			Unique().
			Annotations(entgql.Bind()),
		edge.From("identifiers", PlayerIdentifier.Type).
			Ref("player").
			Annotations(entgql.Bind()),
	}
}
