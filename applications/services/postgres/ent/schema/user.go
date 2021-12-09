package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CheckIn holds the schema definition for the CheckIn entity.
type User struct {
	ent.Schema
}

// Fields of the CheckIn.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MinLen(3).Unique(),
		field.Int64("age").Min(0),
		field.String("password"),
	}
}

// Edges of the CheckIn.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("loans", Loan.Type),
	}
}
