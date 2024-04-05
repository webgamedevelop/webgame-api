package v1

import (
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

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
//	@Success		200	{object}	detailResponse[models.User]
//	@Failure		400	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/user/signup [post]
func (*User) SignUp(c *gin.Context) {
	var (
		user models.User
		err  error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, http.StatusBadRequest, err)
		return
	}

	if _, err = models.CreateUser(&user); err != nil {
		BadResponse(c, http.StatusInternalServerError, err)
		return
	}

	DetailResponse(c, user)
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
//	@Success		200	{object}	detailResponse[models.User]
//	@Failure		400	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
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
			BadResponse(c, http.StatusBadRequest, err)
			return
		}

		username, ok := jwt.ExtractClaims(c)[identityKey]
		if !ok {
			err = fmt.Errorf("key `%s` not found", identityKey)
			BadResponse(c, http.StatusBadRequest, err)
			return
		}

		name, ok := username.(string)
		if !ok {
			err = fmt.Errorf("type assertion failed")
			BadResponse(c, http.StatusBadRequest, err)
			return
		}

		request.Name = name
		if user, err = models.UpdateUser(&request); err != nil {
			BadResponse(c, http.StatusInternalServerError, err)
			return
		}

		DetailResponse(c, user)
		return
	}
}

// ChangePassword change password
//
//	@Tags			user
//	@Summary		change password
//	@Description	change password
//	@Param			user	body	models.UserChangePasswordRequest	true	"change password request"
//	@Produce		json
//	@Success		200	{object}	LoginFailedResponse
//	@Failure		400	{object}	simpleResponse
//	@Failure		401	{object}	simpleResponse
//	@Failure		500	{object}	simpleResponse
//	@Router			/user/password [post]
func (*User) ChangePassword(jwtMiddleware *jwt.GinJWTMiddleware) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			identityKey = jwtMiddleware.IdentityKey
			request     models.UserChangePasswordRequest
			err         error
		)

		if err = c.ShouldBindJSON(&request); err != nil {
			BadResponse(c, http.StatusBadRequest, err)
			return
		}

		username, ok := jwt.ExtractClaims(c)[identityKey]
		if !ok {
			err = fmt.Errorf("key `%s` not found", identityKey)
			BadResponse(c, http.StatusBadRequest, err)
			return
		}

		name, ok := username.(string)
		if !ok {
			err = fmt.Errorf("type assertion failed")
			BadResponse(c, http.StatusBadRequest, err)
			return
		}

		request.Name = name
		if _, err = models.ChangePassword(&request); err != nil {
			BadResponse(c, http.StatusInternalServerError, err)
			return
		}

		jwtMiddleware.LogoutHandler(c)
		return
	}
}
