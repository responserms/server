package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("data", map[string]interface{}{}),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("schema", MetadataSchema.Type).
			Unique(),

		// Back-references to various metadata-supporting entities
		edge.From("user", User.Type).
			Ref("metadata").
			Unique().
			Annotations(entgql.Bind()),
		edge.From("map_type", MapType.Type).
			Ref("metadata").
			Unique().
			Annotations(entgql.Bind()),
	}
}
