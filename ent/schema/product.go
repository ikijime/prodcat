package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
	Attributes []Attribute
}

// Fields of the User.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("code").Unique(),
		field.String("barcode").Unique(),
		field.String("name").MinLen(2),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributeValuesString", AttributeValueString.Type),
		edge.To("attributeValuesNum", AttributeValueNum.Type),
		edge.To("attributeValuesBool", AttributeValueBool.Type),
	}
}
