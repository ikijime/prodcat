package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AttributeVariantNum struct {
	ent.Schema
}

func (AttributeVariantNum) Fields() []ent.Field {
	return []ent.Field{
		field.Int("attribute_id"),
		field.Int("value"),
	}
}

func (AttributeVariantNum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attribute", Attribute.Type).
			Ref("attributeVariantsNum").
			Unique().
			Required().
			Field("attribute_id"),
	}
}
