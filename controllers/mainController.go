package controllers

import (
	"prodcat/services"
	"prodcat/views/layout"

	"github.com/gin-gonic/gin"
)

type MainControler struct {
	authService *services.AuthService
}

func NewMainController(as *services.AuthService) *MainControler {
	return &MainControler{
		authService: as,
	}
}

func (uc MainControler) MainPage(c *gin.Context) {
	c.HTML(200, "", layout.MainMenu("main"))
}
