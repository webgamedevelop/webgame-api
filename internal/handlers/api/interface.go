package api

import "github.com/gin-gonic/gin"

type Webgame interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Detail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type User interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	Update(c *gin.Context)
	Password(c *gin.Context)
}
