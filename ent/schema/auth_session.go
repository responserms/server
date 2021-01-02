package schema

import "github.com/facebook/ent"

// AuthSession holds the schema definition for the AuthSession entity.
type AuthSession struct {
	ent.Schema
}

// Fields of the AuthSession.
func (AuthSession) Fields() []ent.Field {
	return []ent.Field{
		// field.J
	}
}

// Edges of the AuthSession.
func (AuthSession) Edges() []ent.Edge {
	return nil
}
