package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// ServerType holds the schema definition for the ServerType entity.
type ServerType struct {
	ent.Schema
}

// Fields of the ServerType.
func (ServerType) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Unique().
			Immutable(),
		field.String("name"),
		field.Text("description").
			Optional().
			Nillable(),
	}
}

// Edges of the ServerType.
func (ServerType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("servers", Server.Type).
			Ref("server_type").
			Annotations(entgql.Bind()),
	}
}
