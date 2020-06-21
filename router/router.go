package router

import (
	"iv-go/apis"

	"iv-go/pkg"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由设置
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/welcome", apis.Welcome)

	user := router.Group("/users")
	{
		user.GET("/", apis.UserIndex)
	}

	session := router.Group("/", pkg.EnableCookieSession())
	{
		session.POST("/register", apis.Register)
		session.POST("/login", apis.Login)
		session.POST("logout", apis.Logout)

		authorized := session.Group("/", pkg.AuthSessionMiddle())
		{
			authorized.GET("/me", apis.Me)
		}
	}

	return router
}
