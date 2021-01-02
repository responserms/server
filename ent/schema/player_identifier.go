package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// PlayerIdentifier holds the schema definition for the PlayerIdentifier entity.
type PlayerIdentifier struct {
	ent.Schema
}

// Fields of the PlayerIdentifier.
func (PlayerIdentifier) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("identifier"),
	}
}

func (PlayerIdentifier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		MetadataMixin{},
	}
}

// Indexes adds indexes to the fields/edges.
func (PlayerIdentifier) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "identifier").
			Edges("player").
			Unique(),
	}
}

// Edges of the PlayerIdentifier.
func (PlayerIdentifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("player", Player.Type).
			Unique().
			Annotations(entgql.Bind()),
	}
}
