package v1

import (
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"github.com/webgamedevelop/webgame-api/internal/models"
)

// LoginResponse for swagger docs
type LoginResponse struct {
	Code   int    `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

// LoginFailedResponse for swagger docs
type LoginFailedResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User struct{}

// SignUp sign up
//
//	@Tags			user
//	@Summary		sign up
//	@Description	sign up
//	@Param			user	body	models.User	true	"sign up request"
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/user/signup [post]
func (*User) SignUp(c *gin.Context) {
	var (
		user    models.User
		created *models.User
		err     error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		klog.Error(err)
		response := Response(&simpleResponse{}, Code(http.StatusBadRequest), Message(err.Error()))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if created, err = models.CreateUser(&user); err != nil {
		klog.Error(err)
		response := Response(&simpleResponse{}, Code(http.StatusInternalServerError), Message(err.Error()))
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := Response(&simpleResponse{}, Code(http.StatusOK), &ResponseExtend{Data: created})
	c.JSON(http.StatusOK, response)
	return
}

// SignIn sign in
//
//	@Tags			user
//	@Summary		sign in
//	@Description	sign in
//	@Param			user	body		models.UserLoginRequest	true	"login request"
//	@Success		200		{object}	LoginResponse
//	@Failure		401		{object}	LoginFailedResponse
//	@Router			/user/signin [post]
func (*User) SignIn(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context) {
	return jwtMiddleware.LoginHandler
}

// Refresh token
//
//	@Tags			user
//	@Summary		refresh token
//	@Description	refresh token
//	@Success		200	{object}	LoginResponse
//	@Failure		401	{object}	LoginFailedResponse
//	@Router			/user/refresh_token [get]
func (*User) Refresh(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context) {
	return jwtMiddleware.RefreshHandler
}

// SignOut sign out
//
//	@Tags			user
//	@Summary		sign out
//	@Description	sign out
//	@Success		200	{object}	LoginFailedResponse
//	@Router			/user/signout [get]
func (*User) SignOut(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context) {
	return jwtMiddleware.LogoutHandler
}

// Update user info
//
//	@Tags			user
//	@Summary		update user info
//	@Description	update user info
//	@Param			user	body	models.UserUpdateRequest	true	"update user info request"
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/user/update [post]
func (*User) Update(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context) {
	var identityKey = jwtMiddleware.IdentityKey
	return func(c *gin.Context) {
		var (
			request models.UserUpdateRequest
			user    *models.User
			err     error
		)

		if err = c.ShouldBindJSON(&request); err != nil {
			klog.Error(err)
			response := Response(&simpleResponse{}, Code(http.StatusBadRequest), Message(err.Error()))
			c.JSON(http.StatusBadRequest, response)
			return
		}

		username, ok := jwt.ExtractClaims(c)[identityKey]
		if !ok {
			err = fmt.Errorf("key `%s` not found", identityKey)
			klog.Error(err)
			response := Response(&simpleResponse{}, Code(http.StatusBadRequest), Message(err.Error()))
			c.JSON(http.StatusBadRequest, response)
			return
		}

		name, ok := username.(string)
		if !ok {
			err = fmt.Errorf("type assertion failed")
			klog.Error(err)
			response := Response(&simpleResponse{}, Code(http.StatusBadRequest), Message(err.Error()))
			c.JSON(http.StatusBadRequest, response)
			return
		}

		request.Name = name
		if user, err = models.UpdateUser(&request); err != nil {
			klog.Error(err)
			response := Response(&simpleResponse{}, Code(http.StatusInternalServerError), Message(err.Error()))
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		response := Response(&simpleResponse{}, Code(http.StatusOK), &ResponseExtend{Data: user})
		c.JSON(http.StatusOK, response)
		return
	}
}

func (*User) ChangePassword(c *gin.Context) {}
