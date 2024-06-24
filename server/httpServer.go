package server

import (
	"log"
	"prodcat/controllers"
	"prodcat/ent"
	"prodcat/repositories"
	"prodcat/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
}

func InitHttpServer(config *viper.Viper, db *ent.Client) HttpServer {

	authService := services.NewAuthService(db, config)

	// mainControler := controllers.NewMainController(authService)
	userRepo := repositories.NewUserRepository(db, authService)
	productRepo := repositories.NewProductRepository(db)
	// attributeRepo := repositories.NewAttributeRepository(db)
	uploadController := controllers.NewUploadController()
	userController := controllers.NewUsersController(userRepo, authService)
	adminContoller := controllers.NewAdminController(userRepo, authService)
	productController := controllers.NewProductController(productRepo, userRepo, authService)
	attributesController := controllers.NewAttributeController(repositories.NewAttributeRepository(db), authService)
	brandController := controllers.NewBrandController(repositories.NewBrandRepository(db), authService)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(authService.ValidateUser)
	router.HTMLRender = &HTMLTemplRenderer{}

	// authorization
	userOnly := authService.Auth("user")
	adminOnly := authService.Auth("admin")

	router.Static("static", "static")

	router.GET("/", productController.ProductsPage)

	// Authentication
	router.GET("/login", userController.LoginForm)
	router.POST("/login", userController.Login)
	router.GET("/logout", userController.Logout)
	router.GET("/register", userController.RegisterForm)
	router.POST("/register", userController.Register)

	// Uploads
	router.POST("uploads", uploadController.HandleUploads)

	// Products
	router.POST("/products/search", userOnly, productController.SearchProduct)
	router.GET("/products", userOnly, productController.ProductsPage)
	router.GET("/products/:id", userOnly, productController.ProductPage)
	router.GET("/products/:id/edit", userOnly, productController.ProductEditPage)
	router.GET("/products/add", userOnly, productController.ProductAddPage)
	router.POST("/products", userOnly, productController.ProductAddFormParse)
	router.PATCH("/products/:id", userOnly, productController.PatchProductForm)

	router.GET("/api/products", userOnly, productController.GetAllProducts)
	router.POST("/api/products", userOnly, productController.AddProduct)
	router.GET("/api/products/:id", userOnly, productController.GetProductWithAttributeValues)

	// Admin panel
	router.GET("/admin", adminOnly, adminContoller.Index)
	router.GET("/admin/users", adminOnly, adminContoller.Users)
	router.GET("/admin/roles", adminOnly, adminContoller.GetRoles)
	router.GET("/admin/users/:id/edit", adminOnly, userController.UserEdit)
	router.GET("/admin/users/:id", adminOnly, userController.UserGet)
	router.PUT("/admin/users/:id", adminOnly, userController.UserUpdate)
	router.GET("/admin/logs", adminOnly, adminContoller.Logs)

	// Attributes
	router.GET("/attributes", attributesController.GetAllAttributesPage)
	router.POST("/attributes/search", attributesController.SearchAttributes)
	router.GET("/attributes/:id", attributesController.GetAttributePage)
	router.GET("/attributes/:id/edit", attributesController.GetAttributeEditPage)
	router.GET("/attributes/add", attributesController.AddAttributePage)
	router.POST("/attributes", userOnly, attributesController.AttributesAddFormParse)

	router.GET("/attributes/:id/variants/:vid/edit", attributesController.AttributeVariantEdit)
	router.POST("/attributes/:id/variants", attributesController.AttributeVariantAdd)
	router.PUT("/attributes/:id/variants", attributesController.AttributeVariantPut)
	router.DELETE("/attributes/:id/variants/:vid", attributesController.AttributeVariantDelete)

	router.GET("/api/attributes", attributesController.GetAllAttributes)
	router.GET("/api/attributes/:id", attributesController.GetAttribute)
	router.GET("/api/attribute-values", attributesController.GetAllAttributeValues)
	// router.POST("/api/attribute", attributeController.CreateAttribute)

	// Brands
	router.GET("/brands", brandController.GetAllBrandsPage)

	return HttpServer{
		config: config,
		router: router,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
