package controllers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
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

	product := uc.pr.FindProductByID(c, id, true)

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
		log.Println("can't convert")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	product := uc.pr.FindProductByID(c, id, true)

	c.HTML(200, "", fbGuard(c, productView.ProductEdit(dto.NewProductDTO(product))))
}

func (uc ProductController) PatchProductForm(c *gin.Context) {

}

func (uc ProductController) ProductAddFormParse(c *gin.Context) {
	var prod dto.ProductDTO
	c.ShouldBind(&prod)

	image, err := c.FormFile("uploadimage")
	if err != nil {
		log.Println(err.Error())
	} else {
		openedFile, _ := image.Open()
		file, _ := io.ReadAll(openedFile)
		prod.Image.Header = image.Header.Values("Content-Type")[0]
		prod.Image.Name = image.Filename
		prod.Image.Body = base64.RawStdEncoding.EncodeToString(file)
	}

	validate := validator.New()
	err = validate.Struct(&prod)
	if err != nil {
		prod.FillErrors(err)
		c.HTML(http.StatusBadRequest, "", fbGuard(c, productView.ProductAdd(&prod)))
		return
	}

	WritePopupMessage(c, "success", "added")
	emptyProd := dto.ProductDTO{}
	uc.pr.CreateProduct(c, prod)
	c.HTML(http.StatusBadRequest, "", fbGuard(c, productView.ProductAdd(&emptyProd)))
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

	c.JSON(200, uc.pr.FindProductByID(c, id, true))
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
