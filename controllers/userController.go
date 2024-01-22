package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/ent/schema"
	"prodcat/repositories"
	"prodcat/services"
	"prodcat/views"
	"prodcat/views/components"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var cookieMaxAge int = 360 * 24 * 30

type UserController struct {
	userRepo    *repositories.UserRepository
	authService *services.AuthService
}

func NewUsersController(ur *repositories.UserRepository, as *services.AuthService) *UserController {
	return &UserController{
		userRepo:    ur,
		authService: as,
	}
}

func (uc *UserController) LoginForm(c *gin.Context) {
	_, ok := c.Get("user")
	if ok {
		c.Redirect(302, "/")
		return
	}
	userErrors := schema.NewUserErrors()
	c.HTML(200, "", views.Login(c, userErrors))
}

func (uc *UserController) Logout(c *gin.Context) {
	uc.authService.Logout(c)
	c.Redirect(302, "/")
}

func (uc *UserController) Login(c *gin.Context) {
	// Сделать ретраи (возможно с помощью дополнительной таблицы)
	username := c.PostForm("username")
	password := c.PostForm("password")
	userErrors := schema.NewUserErrors()

	// time.Sleep(1 * time.Second)
	user, err := uc.userRepo.FindUserByLogin(c, username)
	if ent.IsNotFound(err) {
		userErrors.UsernameError = fmt.Errorf("user %s not found", username)
		WritePopupMessage(c, "error", "user not found")
		c.HTML(200, "", views.Login(c, userErrors))
		return
	}

	isCorrectPassword := uc.authService.CheckPasswordHash(password, user.Password)
	if !isCorrectPassword {
		userErrors.PasswordError = errors.New("incorrect password or username")
		WritePopupMessage(c, "error", "incorrect password")
		c.HTML(200, "", views.Login(c, userErrors))
		return
	}

	// Login success. Generate JWT token
	signedAccessToken, signedRefreshToken, err := uc.authService.GetTokensForUser(user)
	if err != nil {
		c.HTML(500, "", views.Login(c, nil))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", signedAccessToken, cookieMaxAge, "", "", false, true)
	c.SetCookie("Authorization-ref", signedRefreshToken, cookieMaxAge, "", "", false, true)

	c.Header("HX-Redirect", "/")
}

func (uc *UserController) RegisterForm(c *gin.Context) {
	u, ok := c.Get("user")
	log.Println(u)
	if ok {
		c.Redirect(302, "/")
		return
	}
	userErrors := schema.NewUserErrors()
	c.HTML(200, "", views.Register(c, userErrors))
}

func (uc *UserController) Register(c *gin.Context) {
	// Add more validations
	// Add email or phonenumber confirmation
	var err error

	userErrors := schema.NewUserErrors()

	username := c.PostForm("username")
	password := c.PostForm("password")
	password_confirm := c.PostForm("password_confirm")
	fName := c.PostForm("first_name")
	lName := c.PostForm("last_name")
	email := c.PostForm("email")

	err = schema.ValidateUserUsername(username)
	if err != nil {
		userErrors.UsernameError = errors.New("username not valid")
		WritePopupMessage(c, "error", userErrors.UsernameError.Error())
		c.HTML(http.StatusBadRequest, "", views.Register(c, userErrors))
		return
	}

	_, err = uc.userRepo.FindUserByLogin(c, username)
	if err == nil {
		userErrors.UsernameError = errors.New("user already exists")
		WritePopupMessage(c, "error", userErrors.UsernameError.Error())
		c.HTML(http.StatusBadRequest, "", views.Register(c, userErrors))
		return
	}

	err = schema.ValidateUserPassword(password)

	if err != nil {
		userErrors.PasswordError = errors.New("password not valid")
		WritePopupMessage(c, "error", userErrors.PasswordError.Error()+": "+err.Error())
		c.HTML(http.StatusBadRequest, "", views.Register(c, userErrors))
		return
	}

	if password != password_confirm {
		userErrors.PasswordError = errors.New("passwords don't match")
		WritePopupMessage(c, "error", userErrors.PasswordError.Error())
		c.HTML(http.StatusBadRequest, "", views.Register(c, userErrors))
		return
	}

	userDto := dto.UserDTO{Username: username, Password: password, FirstName: fName, LastName: lName, Email: email}

	user, _ := uc.userRepo.CreateUser(c, userDto)
	WritePopupMessage(c, "success", "User created")

	signedAccessToken, signedRefreshToken, err := uc.authService.GetTokensForUser(user)
	if err != nil {
		c.HTML(500, "", views.Register(c, nil))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", signedAccessToken, cookieMaxAge, "", "", false, true)
	c.SetCookie("Authorization-ref", signedRefreshToken, cookieMaxAge, "", "", false, true)
	c.Header("HX-Redirect", "/")
}

func (uc *UserController) UserEdit(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert")
		c.Redirect(302, "/")
		return
	}

	user, err := uc.userRepo.GetUser(c, id)
	if err != nil {
		log.Panicln(err)
		WritePopupMessage(c, "error", "User not found")
	}

	c.HTML(200, "", components.UserEdit(c, user))
}

func (uc *UserController) UserGet(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		log.Println("can't convert")
		c.Redirect(302, "/")
		return
	}

	user, err := uc.userRepo.GetUser(c, id)
	if err != nil {
		log.Panicln(err)
		WritePopupMessage(c, "error", "User not found")
	}

	c.HTML(200, "", components.UserRow(c, user))
}

func (uc *UserController) UserUpdate(c *gin.Context) {
	type PostParam struct {
		ID          int    `form:"id" json:"id" validate:"required,numeric"`
		FirstName   string `form:"first_name" json:"first_name" validate:"required,min=2,max=15"`
		LastName    string `form:"last_name" json:"last_name" validate:"required,min=2,max=15"`
		Phonenumber string `form:"phonenumber" json:"phonenumber" validate:"omitempty,max=14,numeric"`
		Email       string `form:"email" json:"email"`
		Role        string `form:"role" json:"role" validate:"required"`
	}
	var form PostParam
	var err error
	// var errorString string
	err = c.ShouldBind(&form)
	if err != nil {
		WritePopupMessage(c, "error", "BIND ERROR")
		c.HTML(http.StatusBadRequest, "", components.UserRow(c, nil))
		return
	}

	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		// errorString = fmt.Sprintf("%s %s", "error: ", err.Error())
		WritePopupMessage(c, "error", "INPUT VALIDATION ERROR")
		c.HTML(http.StatusBadRequest, "", components.UserRow(c, nil))
		return
	}

	user, err := uc.userRepo.GetUser(c, form.ID)
	if err != nil {
		WritePopupMessage(c, "success", "User not found")
		c.HTML(http.StatusBadRequest, "", components.UserRow(c, nil))
		return
	}

	updated, err := user.Update().
		SetFirstName(form.FirstName).
		SetLastName(form.LastName).
		SetRole(form.Role).
		SetPhonenumber(form.Phonenumber).
		SetUpdatedAt(time.Now()).
		Save(c)
	if err != nil {
		// errorString = fmt.Sprintf("%s %s", "error: ", err.Error())
		WritePopupMessage(c, "error", "DB VALIDATION ERROR")
		c.HTML(http.StatusBadRequest, "", components.UserRow(c, user))
		return
	}

	WritePopupMessage(c, "success", "User updated successfully")
	c.HTML(200, "", components.UserRow(c, updated))
}
