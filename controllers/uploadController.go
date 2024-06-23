package controllers

import (
	"github.com/gin-gonic/gin"
)

type UploadController struct {
}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (uc UploadController) HandleUploads(c *gin.Context) {
	formData, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(500, err)
	}
	_ = formData
	c.Status(200)
}
