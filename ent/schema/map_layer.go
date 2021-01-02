package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// MapLayer holds the schema definition for the MapLayer entity.
type MapLayer struct {
	ent.Schema
}

// Fields of the MapLayer.
func (MapLayer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("url_template"),
		field.Bool("is_public"),
	}
}

// Mixin injects additional fields via mixins.
func (MapLayer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		MetadataMixin{},
	}
}

// Edges of the MapLayer.
func (MapLayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("map_type", MapType.Type).
			Annotations(entgql.Bind()).
			Unique(),
	}
}
