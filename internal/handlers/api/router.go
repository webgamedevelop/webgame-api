//	@title			webgame-api
//	@version		v1
//	@description	webgame-api docs
//	@contact.name	webgamedevelop
//	@contact.email	webgamedevelop@163.com
//	@contact.url	http://www.swagger.io/support
//	@host			localhost:8080
//	@BasePath		/api/v1

package api

import (
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AttachUserAPI(group *gin.RouterGroup, user User, jwtMiddleware *jwt.GinJWTMiddleware) {
	userAPI := group.Group("/user")
	userAPI.POST("/signup", user.SignUp)
	userAPI.POST("/signin", user.SignIn(jwtMiddleware))
	userAPI.GET("/refresh_token", user.Refresh(jwtMiddleware))
	userAPI.GET("/signout", user.SignOut(jwtMiddleware))
	userAPI.Use(jwtMiddleware.MiddlewareFunc())
	userAPI.POST("/update", user.Update(jwtMiddleware))
	userAPI.POST("/password", user.ChangePassword(jwtMiddleware))
}

func AttachResourceAPI(group *gin.RouterGroup, root string, resource Resource) {
	if !strings.HasPrefix(root, "/") {
		panic("root must begin with '/'")
	}
	api := group.Group(root)
	api.POST("/create", resource.Create)
	api.POST("/update", resource.Update)
	api.GET("/list", resource.List)
	api.GET("/detail", resource.Detail)
	api.DELETE("/delete", resource.Delete)
	api.GET("/syncto", resource.SyncTo)
	api.GET("/syncfrom", resource.SyncFrom)
}
