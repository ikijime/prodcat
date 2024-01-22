package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type UserSettings struct {
	ent.Schema
	Attributes []Attribute
}

type DefaultSetting struct {
	ProductView string `json:"productViewType"`
}

func getDefaults() DefaultSetting {
	return DefaultSetting{ProductView: "rows"}
}

// Fields of the User.
func (UserSettings) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("frontend", DefaultSetting{}).Default(getDefaults()),
	}
}

func (UserSettings) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("user", User.Type),
	}
}
