package v1

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) SignIn(c *gin.Context) {}

func (*User) SignUp(c *gin.Context) {}

func (*User) SignOut(c *gin.Context) {}

func (*User) Update(c *gin.Context) {}

func (*User) Password(c *gin.Context) {}
