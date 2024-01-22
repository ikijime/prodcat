package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AttributeVariantString struct {
	ent.Schema
}

func (AttributeVariantString) Fields() []ent.Field {
	return []ent.Field{
		field.Int("attribute_id").Unique(),
		field.String("value").Unique().MinLen(1),
	}
}

func (AttributeVariantString) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attribute", Attribute.Type).
			Ref("attributeVariantsString").
			Unique().
			Required().
			Field("attribute_id"),
	}
}
