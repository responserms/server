package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("api_username").
			Optional().
			Nillable(),
		field.String("api_secret").
			Optional().
			Nillable(),
		field.String("api_address").
			Optional().
			Nillable(),
		field.String("api_port").
			Optional().
			Nillable(),
	}
}

func (Server) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		MetadataMixin{},
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server_type", ServerType.Type).
			Unique().
			Annotations(entgql.Bind()),
		edge.To("map_type", MapType.Type).
			Unique().
			Annotations(entgql.Bind()),
		edge.From("players", Player.Type).
			Ref("servers").
			Annotations(entgql.Bind()),
	}
}
