package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Loan holds the schema definition for the Loan entity.
type Loan struct {
	ent.Schema
}

// Fields of the Loan.
func (Loan) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("entityId"),
		field.Enum("status").
			Values("RESERVED", "PICKED", "PACKED", "SHIPPED"),
	}
}

// Edges of the Loan.
func (Loan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("loans").
			Unique(),
	}
}
