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
	userAPI.POST("/update", user.Update)
	userAPI.POST("/password", user.ChangePassword)
}

func AttachWebgameAPI(group *gin.RouterGroup, webgame Webgame) {
	webgameAPI := group.Group("/webgame")
	webgameAPI.POST("/create", webgame.Create)
	webgameAPI.GET("/list", webgame.List)
	webgameAPI.GET("/detail", webgame.Detail)
	webgameAPI.POST("/update", webgame.Update)
	webgameAPI.DELETE("/delete", webgame.Delete)
}
