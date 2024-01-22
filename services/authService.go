package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"prodcat/ent"
	"prodcat/views"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var CookieExpireTime int = 360 * 24 * 30

var TokenExpireTime = func() int64 {
	loc, _ := time.LoadLocation(os.Getenv("TZ"))
	return time.Now().In(loc).Add(time.Minute * 1).Unix()
}

var RefreshTokenExpireTime = func() int64 {
	loc, _ := time.LoadLocation(os.Getenv("TZ"))
	return time.Now().In(loc).Add(time.Hour * 24).Unix()
}

// todo db without repo, not good
type AuthService struct {
	secret_key string
	db         *ent.Client
	config     *viper.Viper
}

type UserRefClaims struct {
	ID int
	jwt.StandardClaims
}

type UserClaims struct {
	ID        int
	FirstName string
	LastName  string
	UserRoles string
	jwt.StandardClaims
}

func NewAuthService(db *ent.Client, config *viper.Viper) *AuthService {
	return &AuthService{
		secret_key: config.GetString("jwt.token_secret"),
		db:         db,
		config:     config,
	}
}

func (as AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (as AuthService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ------------------- JWT -------------------------------------------------
func (au AuthService) NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(au.secret_key))
}

func (au AuthService) NewRefreshToken(claims UserRefClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(au.secret_key))
}

func (au AuthService) ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(au.secret_key), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func (au AuthService) ParseRefreshToken(refreshToken string) *UserRefClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &UserRefClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(au.secret_key), nil
	})

	return parsedRefreshToken.Claims.(*UserRefClaims)
}

func (au AuthService) GetTokensForUser(user *ent.User) (string, string, error) {
	userClaims := UserClaims{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserRoles: user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: TokenExpireTime(),
		},
	}

	signedAccessToken, err := au.NewAccessToken(userClaims)
	if err != nil {
		log.Fatal("error creating access token")
	}

	refreshClaims := UserRefClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: RefreshTokenExpireTime(),
		},
	}

	signedRefreshToken, err := au.NewRefreshToken(refreshClaims)
	if err != nil {
		log.Fatal("error creating refresh token")
	}

	return signedAccessToken, signedRefreshToken, nil
}

// ------------------- Middlewares -------------------------------------------------
func (au AuthService) ValidateUser(c *gin.Context) {
	tokenStr, refreshTokenStr, err := au.checkCookieTokens(c)
	if err != nil {
		log.Println("no tokens. continue...", err)
		c.Next()
		return
	}

	refreshToken, err := au.validateToken(refreshTokenStr)
	if err != nil {
		log.Println(err)
		c.Next()
		return
	}

	claims, _ := refreshToken.Claims.(jwt.MapClaims)
	id, _ := claims["ID"].(float64)
	user, err := au.getUser(c, int(id))
	if err != nil {
		log.Println(err)
		c.Next()
		return
	}

	token, err := au.validateToken(tokenStr)

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			err = nil
			newTokenString, newRefreshTokenString, err := au.refreshTokens(user, token, refreshToken)
			if err != nil {
				log.Println("FAILED TO REFRESH TOKENS")
				log.Println(err)
				c.Next()
				return
			}
			log.Println("Refreshing tokens")
			c.SetCookie("Authorization", newTokenString, CookieExpireTime, "", "", false, true)
			c.SetCookie("Authorization-ref", newRefreshTokenString, CookieExpireTime, "", "", false, true)
		}
	}

	if err != nil {
		log.Println(err)
		c.Next()
		return
	}

	IsLoggedIn := false
	IsAdmin := false

	if user != nil {
		IsLoggedIn = true
		roles := user.Role
		if roles == "admin" {
			IsAdmin = true
		}
		// Global context variables
		c.Set("isLoggedIn", IsLoggedIn)
		c.Set("isAdmin", IsAdmin)
	}

	c.Set("user", user)
	st := user.QuerySettings().AllX(c)
	c.Set("user_settings", st)
	c.Next()
}

func (au AuthService) Auth(role string) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		user, ok := c.Get("user")

		if !ok {
			fmt.Println("no user")
			c.HTML(403, "", views.Forbidden())
			c.Abort()
			return
		}

		userInfered, ok := user.(*ent.User)
		if !ok {
			fmt.Println("cannot infer user")
			c.HTML(403, "", views.Forbidden())
			c.Abort()
			return
		}

		userRole := userInfered.Role
		rolesNeeded := au.GetAvailableRoles(c)[userRole]

		// fmt.Println(userRole)
		// fmt.Println(rolesNeeded)
		canAccess := slices.Contains(rolesNeeded, role)
		if !canAccess {
			fmt.Println("no access")
			c.HTML(403, "", views.Forbidden())
			c.Abort()
			return
		}

		c.Next()
	}

	return gin.HandlerFunc(fn)
}

func (au AuthService) AuthAPI(role string) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		user, ok := c.Get("user")

		if !ok {
			c.JSON(403, "")
			c.Abort()
			return
		}

		userInfered, ok := user.(*ent.User)
		if !ok {
			c.JSON(403, "")
			c.Abort()
			return
		}

		userRole := userInfered.Role
		rolesNeeded := au.GetAvailableRoles(c)[userRole]

		canAccess := slices.Contains(rolesNeeded, role)
		if !canAccess {
			c.JSON(403, "")
			c.Abort()
			return
		}

		c.Next()
	}

	return gin.HandlerFunc(fn)
}

// ---------------------------------------------------------------------------------

func (au AuthService) checkCookieTokens(c *gin.Context) (string, string, error) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil || tokenString == "" {
		return "", "", errors.New("no token string")
	}

	refreshTokenString, err := c.Cookie("Authorization-ref")

	if err != nil || refreshTokenString == "" {
		return "", "", errors.New("no refresh token string")
	}

	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func (au AuthService) validateToken(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}

		return []byte(au.secret_key), nil
	})

	if err != nil {
		return nil, err
	}

	err = token.Claims.Valid()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (au AuthService) refreshTokens(u *ent.User, token *jwt.Token, refreshToken *jwt.Token) (string, string, error) {

	signedAccessToken, signedRefreshToken, err := au.GetTokensForUser(u)
	if err != nil {
		return "", "", fmt.Errorf("couldn't generate user tokents, user %d: ", u.ID)
	}

	return signedAccessToken, signedRefreshToken, nil
}

func (au AuthService) getUser(c *gin.Context, id int) (*ent.User, error) {
	user, err := au.db.User.Get(c, int(id))
	if err != nil {
		return nil, fmt.Errorf("user %d not found: ", id)
	}
	return user, nil
}

func (au AuthService) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 1, "", "", false, true)
	c.SetCookie("Authorization-ref", "", 1, "", "", false, true)
}

func (au AuthService) GetAvailableRoles(c *gin.Context) map[string][]string {
	return au.config.GetStringMapStringSlice("roles")
}
