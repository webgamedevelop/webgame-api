package middleware

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/webgamedevelop/webgame-api/internal/models"
)

var IdentityKey = "name"

// JWT init jwt middleware
func JWT() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "webgame-api realm",
		Key:           []byte("webgame-api secret key"),
		Timeout:       timeout,
		MaxRefresh:    maxRefresh,
		IdentityKey:   IdentityKey,
		TokenLookup:   fmt.Sprintf("header: Authorization, query: %s, cookie: %s", cookieName, cookieName),
		TokenHeadName: "Bearer",
		SendCookie:    true,
		CookieName:    cookieName,
		CookieMaxAge:  timeout,
		TimeFunc:      time.Now,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginUser models.LoginUser
			if err := c.ShouldBind(&loginUser); err != nil {
				return "", errors.Join(err, jwt.ErrMissingLoginValues)
			}
			if err := models.CompareUser(loginUser.Name, loginUser.Password); err != nil {
				return nil, errors.Join(err, jwt.ErrFailedAuthentication)
			}
			return &models.User{Name: loginUser.Name}, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{IdentityKey: v.Name}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{Name: claims[IdentityKey].(string)}
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(code, gin.H{"code": code, "message": "ok"})
		},
	})

	if err != nil {
		return nil, err
	}

	if err := authMiddleware.MiddlewareInit(); err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
