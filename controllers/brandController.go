package controllers

import (
	"prodcat/ent/schema"
	"prodcat/repositories"
	"prodcat/services"
	"prodcat/utils"
	brandView "prodcat/views/brand"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandController struct {
	br          *repositories.BrandRepository
	authService *services.AuthService
}

func NewBrandController(br *repositories.BrandRepository, as *services.AuthService) *BrandController {
	return &BrandController{
		br:          br,
		authService: as,
	}
}

// ///////// ATRIBUTES /////////////////////////////////////////
func (bc BrandController) GetAllBrandsPage(c *gin.Context) {
	htmxHeader := c.GetHeader("HX-Request")
	fallback := false
	if htmxHeader == "" {
		fallback = true
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	aType := c.Query("type")
	if limit == 0 {
		limit = 10
		offset = 0
	}

	if !utils.StringInSlice(aType, schema.AttrTypes) {
		aType = ""
	}

	brands := bc.br.GetAllBrands(c, limit, offset)
	c.HTML(200, "", fbGuard(c, brandView.Brands(c, brands, fallback)))
}

// func (ac AttributeController) GetAllBrands(c *gin.Context) {
// 	idQuery := c.Param("id")
// 	id, err := strconv.Atoi(idQuery)
// 	if err != nil {
// 		log.Println("can't convert id")
// 		return
// 	}
// 	attribute := ac.ar.GetAttributeWithVariants(c, id)
// 	c.HTML(200, "", views.Attribute(c, dto.NewAttrWithVariants(attribute)))
// }
