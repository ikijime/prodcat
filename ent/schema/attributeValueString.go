package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AttributeValueString struct {
	ent.Schema
}

func (AttributeValueString) Fields() []ent.Field {
	return []ent.Field{
		field.Int("variant_id"),
		field.Int("product_id"),
	}
}

func (AttributeValueString) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("variant", AttributeVariantString.Type).
			Unique().
			Required().
			Field("variant_id"),
		edge.From("product", Product.Type).
			Ref("attributeValuesString").
			Required().
			Unique().
			Field("product_id"),
	}
}

func (AttributeValueString) IsAttributeValue() bool {
	return true
}
