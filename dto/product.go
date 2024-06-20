package dto

import (
	"errors"
	"prodcat/ent"
	ut "prodcat/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ProductDTO struct {
	ID             int
	Name           string `form:"name" json:"name" validate:"required,min=2"`
	NameErr        error
	Code           string `form:"code" json:"code" validate:"required,numeric,min=1"`
	CodeErr        error
	Barcode        string `form:"barcode" json:"barcode" validate:"required,min=2"`
	BarcodeErr     error
	Description    string `form:"description" json:"description"`
	DescriptionErr error
	Err            error
	ImageHeaders   string
	Image          Image
}

func (d *ProductDTO) IsValid() bool {
	return false
}

func NewProductDTO(p *ent.Product) ProductDTO {
	return ProductDTO{ID: p.ID, Name: p.Name, Code: strconv.Itoa(p.Code), Barcode: p.Barcode, Description: p.Description}
}

func (d *ProductDTO) FillErrors(err error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch field {
			case "Code":
				d.CodeErr = ut.MsgForTag(err)
			case "Barcode":
				d.BarcodeErr = ut.MsgForTag(err)
			case "Name":
				d.NameErr = ut.MsgForTag(err)
			default:
				d.Err = ut.MsgForTag(err)
			}
		}
	}
}
