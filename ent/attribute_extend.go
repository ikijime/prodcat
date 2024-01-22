package ent

type AttributeVariantI interface {
	AttributeOrErr() (*Attribute, error)
}

// func (a *Attribute) GetVariants() []*AttributeVariantI {
// 	switch a.Type {
// 	case "string":
// 		return a.Edges.AttributeVariantsString
// 	default:
// 		panic("attribute type not correct")
// 	}
// }
