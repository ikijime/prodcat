package dto

import (
	"errors"
	"prodcat/ent"
	ut "prodcat/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type AttributeValue interface {
	GetValue(string) (ent.Value, error)
}

type AttributeCollection struct {
	BoolAttrs   []AttributeDTO
	StringAttrs []AttributeDTO
	NumAttrs    []AttributeDTO
}

type AttributeDTO struct {
	ID             int    `form:"attr_id" json:"id"`
	Type           string `form:"type" json:"type" validate:"required,min=2"`
	TypeErr        error
	Name           string `form:"name" json:"name" validate:"required,min=2"`
	NameErr        error
	Description    string `form:"description" json:"description" validate:"required,min=2"`
	DescriptionErr error
	ValueID        int
	Value          string
	Err            error
}

func NewAttributeDTO(attr *ent.Attribute) AttributeDTO {
	return AttributeDTO{ID: attr.ID, Type: attr.Type, Name: attr.Name, Description: attr.Description}
}

func (d *AttributeDTO) FillErrors(err error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch field {
			case "Description":
				d.DescriptionErr = ut.MsgForTag(err)
			case "Name":
				d.NameErr = ut.MsgForTag(err)
			case "Type":
				d.TypeErr = ut.MsgForTag(err)
			default:
				d.Err = ut.MsgForTag(err)
			}
		}
	}
}

type AttributeWithVariantsDTO struct {
	A AttributeDTO
	V []AttributeVariantDTO
}

type AttributeVariantDTO struct {
	ID          int    `form:"var_id" json:"id"`
	AttributeID int    `form:"attr_id" json:"attr_id"`
	Value       string `form:"value" json:"value" validate:"min=1"`
	ValueErr    error
	Type        string `form:"type" json:"type required"`
	Err         error
}

func (d *AttributeVariantDTO) FillErrors(err error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch field {
			case "Value":
				d.ValueErr = ut.MsgForTag(err)
			default:
				d.Err = ut.MsgForTag(err)
			}
		}
	}
}

func CheckAttributeType(atype string) bool {
	switch atype {
	case "string":
		return true
	case "numeric":
		return true
	case "bool":
		return true
	default:
		return false
	}
}

func NewAttrWithVariants(a *ent.Attribute) AttributeWithVariantsDTO {
	var variantsDTOs []AttributeVariantDTO
	switch a.Type {
	case "string":
		for _, v := range a.Edges.AttributeVariantsString {
			variantsDTOs = append(variantsDTOs, AttributeVariantDTO{
				ID:          v.ID,
				AttributeID: a.ID,
				Type:        "string",
				Value:       v.Value,
			})
		}
	case "numeric":
		for _, v := range a.Edges.AttributeVariantsNum {
			variantsDTOs = append(variantsDTOs, AttributeVariantDTO{
				ID:          v.ID,
				AttributeID: a.ID,
				Type:        "numeric",
				Value:       strconv.Itoa(v.Value),
			})
		}
	case "bool":

	default:
		panic("not valid attribute type: " + a.Type)
	}

	return AttributeWithVariantsDTO{
		A: AttributeDTO{Name: a.Name, Description: a.Description, ID: a.ID, Type: a.Type},
		V: variantsDTOs,
	}
}

func NewAttributeCollection(product *ent.Product) AttributeCollection {
	var sAttributes []AttributeDTO
	var bAttributes []AttributeDTO
	var nAttributes []AttributeDTO
	for _, v := range product.Edges.AttributeValuesString {
		newAttr := AttributeDTO{
			ID:          v.Edges.Variant.Edges.Attribute.ID,
			Type:        v.Edges.Variant.Edges.Attribute.Type,
			Name:        v.Edges.Variant.Edges.Attribute.Name,
			Description: v.Edges.Variant.Edges.Attribute.Description,
			ValueID:     v.Edges.Variant.ID,
			Value:       v.Edges.Variant.Value,
		}
		sAttributes = append(sAttributes, newAttr)
	}

	for _, v := range product.Edges.AttributeValuesBool {
		newAttr := AttributeDTO{
			ID:          v.Edges.Attribute.ID,
			Type:        v.Edges.Attribute.Type,
			Name:        v.Edges.Attribute.Name,
			Description: v.Edges.Attribute.Description,
			ValueID:     v.ID,
			Value:       strconv.FormatBool(v.Value),
		}
		bAttributes = append(bAttributes, newAttr)
	}

	for _, v := range product.Edges.AttributeValuesNum {
		newAttr := AttributeDTO{
			ID:          v.Edges.Variant.Edges.Attribute.ID,
			Type:        v.Edges.Variant.Edges.Attribute.Type,
			Name:        v.Edges.Variant.Edges.Attribute.Name,
			Description: v.Edges.Variant.Edges.Attribute.Description,
			ValueID:     v.Edges.Variant.ID,
			Value:       strconv.Itoa(v.Edges.Variant.Value),
		}
		nAttributes = append(nAttributes, newAttr)
	}

	return AttributeCollection{BoolAttrs: bAttributes, StringAttrs: sAttributes, NumAttrs: nAttributes}
}
