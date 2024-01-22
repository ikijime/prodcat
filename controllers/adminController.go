package controllers

import (
	"log"
	"prodcat/repositories"
	"prodcat/services"
	adminViews "prodcat/views/admin"
	"prodcat/views/components"
	"prodcat/views/layout"

	"github.com/gin-gonic/gin"
)

type AdminContoller struct {
	userRepo    *repositories.UserRepository
	authService *services.AuthService
}

func NewAdminController(ur *repositories.UserRepository, as *services.AuthService) *AdminContoller {
	return &AdminContoller{
		userRepo:    ur,
		authService: as,
	}
}

func (ac AdminContoller) Index(c *gin.Context) {
	c.HTML(200, "", layout.AdminMenu(c))
}

func (ac AdminContoller) Users(c *gin.Context) {
	userlist, err := ac.userRepo.GetUsers(c, 0, 10)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.HTML(200, "", adminViews.AdminUsers(c, userlist))
}

func (ac AdminContoller) GetRoles(c *gin.Context) {
	roles := ac.authService.GetAvailableRoles(c)

	keys := make([]string, len(roles))

	i := 0
	for k := range roles {
		keys[i] = k
		i++
	}

	c.HTML(200, "", components.Roles(c, keys))
}

func (ac AdminContoller) ChangeRoleModal(c *gin.Context) {
	id := c.Query("id")
	log.Println(id)
	roles := ac.authService.GetAvailableRoles(c)

	keys := make([]string, len(roles))

	i := 0
	for k := range roles {
		keys[i] = k
		i++
	}

	c.HTML(200, "", components.ModalRole(c, keys, "user"))
}

func (ac AdminContoller) Logs(c *gin.Context) {
	c.HTML(200, "", adminViews.AdminLogs(c))
}
