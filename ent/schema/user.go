package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/responserms/server/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("password").
			Optional().
			Sensitive().
			Nillable(),
		field.Bool("is_system").
			Default(false),
		field.Time("disabled_at").
			Optional().
			Nillable(),
		field.Text("disabled_reason").
			Optional().
			Nillable(),
		field.Time("activated_at").
			Optional().
			Nillable(),
		field.Text("activation_comment").
			Optional().
			Nillable(),
	}
}

// Mixin injects the mixins.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MetadataMixin{},
		mixin.Time{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("session_tokens", SessionToken.Type).
			Annotations(entgql.Bind()),
		edge.From("activation", Activation.Type).
			Ref("user").
			Unique().
			Annotations(entgql.Bind()),
		edge.From("activations", Activation.Type).
			Ref("actor").
			Annotations(entgql.Bind()),
		edge.From("players", Player.Type).
			Ref("user").
			Annotations(entgql.Bind()),
	}
}
