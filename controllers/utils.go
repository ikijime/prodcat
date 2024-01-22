package controllers

import (
	"prodcat/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func WritePopupMessage(c *gin.Context, typeOf string, message string) {
	c.Writer.Header().Set("HX-Trigger", "{\"showMessage\": {\"message\":\""+message+"\", \"messageType\": \""+typeOf+"\"}}")
}

func fbGuard(c *gin.Context, cmp templ.Component) templ.Component {
	fallback := false
	if c.GetHeader("HX-Request") == "" {
		fallback = true
	}
	return views.Fallback(c, cmp, fallback)
}
