package repositories

import (
	"net/http/httptest"
	"testing"

	"prodcat/dto"
	"prodcat/ent/enttest"
	"prodcat/ent/schema"

	"github.com/gin-gonic/gin"
	_ "github.com/xiaoqidun/entps"
)

func TestCreateProductWithAttributes(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	productRepo := NewProductRepository(client)
	attributeRepo := NewAttributeRepository(client)

	// attributeRepo := repositories.New
	name := "NewProduct"
	description := "Some description"
	product, err := productRepo.CreateProduct(c, dto.ProductDTO{Name: name, Description: description})

	if err != nil {
		t.Errorf("cant' create new product")
	}

	attrBool := attributeRepo.CreateAttribute(c, dto.AttributeDTO{Type: schema.AttrBoolType, Name: "boolAttr", Description: "it is boolean"})
	attrString := attributeRepo.CreateAttribute(c, dto.AttributeDTO{Type: schema.AttrStringType, Name: "stringAttr", Description: "it is string"})
	attrString2 := attributeRepo.CreateAttribute(c, dto.AttributeDTO{Type: schema.AttrStringType, Name: "stringAttr", Description: "it is string2"})

	attrVARIANT, _ := attributeRepo.CreateStringVariant(c, attrString, "attribute value")
	attrVARIANT2, _ := attributeRepo.CreateStringVariant(c, attrString2, "attribute2 value")
	attributeRepo.AttachVariantToProduct(c, product, attrVARIANT)
	attributeRepo.AttachVariantToProduct(c, product, attrVARIANT2)

	attributeRepo.CreateBoolValue(c, attrBool, false, product)

	updatedProd := productRepo.FindProductByID(c, product.ID, true)

	boolAttrs := updatedProd.Edges.AttributeValuesBool
	strAttrs := updatedProd.Edges.AttributeValuesString
	println(boolAttrs[0].Value)
	println(strAttrs)
}
