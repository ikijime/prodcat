package repositories

import (
	"net/http/httptest"
	"testing"

	"prodcat/dto"
	"prodcat/ent/enttest"

	"github.com/gin-gonic/gin"
	_ "github.com/xiaoqidun/entps"
)

func TestWriteProduct(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	productRepo := NewProductRepository(client)
	name := "NewProduct"
	description := "Some description"
	newProduct, err := productRepo.CreateProduct(c, dto.ProductDTO{Name: name, Description: description})
	if err != nil {
		t.Errorf("cant' create new product")
	}

	if newProduct.Name != name && newProduct.Description != description {
		t.Errorf("products fields don't match")
	}
}
