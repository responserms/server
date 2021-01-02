package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// MapType holds the schema definition for the MapType entity.
type MapType struct {
	ent.Schema
}

// Fields of the MapType.
func (MapType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("min_zoom"),
		field.Int("max_zoom"),
		field.Float("min_x"),
		field.Float("min_y"),
		field.Float("max_x"),
		field.Float("max_y"),
	}
}

func (MapType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		MetadataMixin{},
	}
}

// Edges of the MapType.
func (MapType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("map_layers", MapLayer.Type).
			Ref("map_type").
			Annotations(entgql.Bind()),
		edge.From("servers", Server.Type).
			Ref("map_type").
			Annotations(entgql.Bind()),
	}
}
