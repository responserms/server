package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/qri-io/jsonschema"
	"github.com/responserms/server/ent/schema/mixin"
)

// MetadataSchema holds the schema definition for the MetadataSchema entity.
//
// The MetadataSchema entity holds a JSON Schema object used to validate a metadata
// object. Because metadata is an external-use feature and not something that Response
// uses internally, by default without a schema no validation is performed. This schema
// allows Response integrators to validate metadata added to their entity records.
type MetadataSchema struct {
	ent.Schema
}

// Fields of the MetadataSchema.
func (MetadataSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Text("about").
			Optional().
			Nillable(),
		field.JSON("schema", &jsonschema.Schema{}),
	}
}

// Mixin registeres mixins that add additional data to the schema.
func (MetadataSchema) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the MetadataSchema.
func (MetadataSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).
			Ref("schema"),
	}
}
