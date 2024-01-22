package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AttributeValueBool struct {
	ent.Schema
}

func (AttributeValueBool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("attribute_id"),
		field.Int("product_id"),
		field.Bool("value"),
	}
}

func (AttributeValueBool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attribute", Attribute.Type).
			Ref("attributeValuesBool").
			Unique().
			Required().Field("attribute_id"),
		edge.From("product", Product.Type).
			Ref("attributeValuesBool").
			Unique().
			Required().Field("product_id"),
	}
}

func (AttributeValueBool) IsAttributeValue() bool {
	return true
}
