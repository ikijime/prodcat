package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AttributeValueNum struct {
	ent.Schema
}

func (AttributeValueNum) Fields() []ent.Field {
	return []ent.Field{
		field.Int("variant_id"),
		field.Int("product_id"),
	}
}

func (AttributeValueNum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("variant", AttributeVariantNum.Type).
			Unique().
			Required().
			Field("variant_id"),
		edge.From("product", Product.Type).
			Ref("attributeValuesNum").
			Required().
			Unique().
			Field("product_id"),
	}
}

func (AttributeValueNum) IsAttributeValue() bool {
	return true
}
