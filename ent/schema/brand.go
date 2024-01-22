package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Brand struct {
	ent.Schema
}

// Fields of the User.
func (Brand) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Brand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
	}
}
