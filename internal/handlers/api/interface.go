package api

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Webgame interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Detail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type User interface {
	SignUp(c *gin.Context)
	SignIn(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context)
	Refresh(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context)
	SignOut(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context)
	Update(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context)
	ChangePassword(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context)
}
