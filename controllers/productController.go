package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/repositories"
	"prodcat/services"
	productView "prodcat/views/product"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	pr          *repositories.ProductRepository
	ur          *repositories.UserRepository
	authService *services.AuthService
}

func NewProductController(pr *repositories.ProductRepository, ur *repositories.UserRepository, as *services.AuthService) *ProductController {
	return &ProductController{
		pr:          pr,
		ur:          ur,
		authService: as,
	}
}

func (uc ProductController) ProductsPage(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	if limit == 0 {
		limit = 10
		offset = 0
	}

	viewMode := "rows"
	viewModeQuery := c.Query("view")

	if viewModeQuery != "" {
		viewMode = viewModeQuery
	}

	uSettings, ok := c.Get("user_settings")
	if ok {
		userSettings := uSettings.([]*ent.UserSettings)[0]
		viewMode = userSettings.Frontend.ProductView

		if viewModeQuery != "" {
			userSettings.Frontend.ProductView = viewModeQuery
			userSettings.Update().SetFrontend(userSettings.Frontend).SaveX(c)
			viewMode = viewModeQuery
		}
	}

	searchQuery := c.Query("search")
	if searchQuery != "" {
		products := uc.pr.SearchProduct(c, offset, limit, searchQuery)
		c.HTML(200, "", fbGuard(c, productView.Products(products, viewMode)))
		return
	}
	c.HTML(200, "", fbGuard(c, productView.Products(uc.pr.GetAllProducts(c, limit, offset), viewMode)))
}

func (uc ProductController) ProductPage(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	product, err := uc.pr.FindProductByID(c, id, true)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	attributeCollection := dto.NewAttributeCollection(product)
	c.HTML(200, "", fbGuard(c, productView.Product(product, attributeCollection)))
}

func (uc ProductController) ProductAddPage(c *gin.Context) {
	productDTO := dto.ProductDTO{}
	c.HTML(200, "", fbGuard(c, productView.ProductAdd(&productDTO)))
}

func (uc ProductController) ProductEditPage(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.HTML(404, "can't convert", "")
		return
	}

	product, err := uc.pr.FindProductByID(c, id, true)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.HTML(200, "", fbGuard(c, productView.ProductEdit(dto.NewProductDTO(product))))
}

func (uc ProductController) PatchProductForm(ctx *gin.Context) {
	var prodDTO dto.ProductDTO
	err := ctx.ShouldBind(&prodDTO)
	if err != nil {
		ctx.JSON(200, err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(&prodDTO)
	if err != nil {
		prodDTO.FillErrors(err)
		ctx.HTML(http.StatusBadRequest, "", fbGuard(ctx, productView.ProductEdit(prodDTO)))
		return
	}

	product, _ := uc.pr.UpdateProduct(ctx, &prodDTO)
	ctx.HTML(200, "", productView.ProductEdit(dto.NewProductDTO(product)))
}

func (uc ProductController) ProductAddFormParse(c *gin.Context) {
	var prod dto.ProductDTO
	c.ShouldBind(&prod)

	validate := validator.New()
	err := validate.Struct(&prod)
	if err != nil {
		prod.FillErrors(err)
		c.HTML(422, "", fbGuard(c, productView.ProductAddForm(&prod)))
		return
	}

	form, err := c.MultipartForm()
	if err == nil {
		for _, handlers := range form.File {
			for _, handler := range handlers {
				dst, _ := os.Create(handler.Filename)
				file, _ := handler.Open()
				defer dst.Close()
				defer file.Close()
				if _, err := io.Copy(dst, file); err != nil {
					c.AbortWithError(500, err)
					return
				}
			}
		}
	}

	WritePopupMessage(c, "success", "added")
	emptyProd := dto.ProductDTO{}
	uc.pr.CreateProduct(c, prod)
	c.HTML(201, "", fbGuard(c, productView.ProductAddForm(&emptyProd)))
}

func (uc ProductController) SearchProduct(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	if limit == 0 {
		limit = 1
		offset = 0
	}

	viewMode := "rows"
	uSettings, ok := c.Get("user_settings")
	if ok {
		userSettings := uSettings.([]*ent.UserSettings)[0]
		viewMode = userSettings.Frontend.ProductView
	}
	c.HTML(200, "", productView.Products(uc.pr.GetAllProducts(c, limit, offset), viewMode))
}

func (uc ProductController) GetAllProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	if limit == 0 {
		limit = 10
		offset = 0
	}
	c.JSON(200, uc.pr.GetAllProducts(c, limit, offset))
}

func (uc ProductController) GetProductWithAttributeValues(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert")
		return
	}
	product, _ := uc.pr.FindProductByID(c, id, true)
	c.JSON(200, product)
}

func (uc ProductController) AddProduct(c *gin.Context) {
	var prod dto.ProductDTO
	err := c.ShouldBind(&prod)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	product, err := uc.pr.CreateProduct(c, prod)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": product})
}
