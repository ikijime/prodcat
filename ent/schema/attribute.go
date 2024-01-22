package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const AttrStringType = "string"
const AttrBoolType = "bool"
const AttrNumType = "numeric"

var AttrTypes = []string{"string", "bool", "numeric"}

type Attribute struct {
	ent.Schema
}

func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().MinLen(2),
		field.String("description"),
		field.String("type"),
	}
}

func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributeValuesBool", AttributeValueBool.Type),
		edge.To("attributeVariantsString", AttributeVariantString.Type),
		edge.To("attributeVariantsNum", AttributeVariantNum.Type),
	}
}
