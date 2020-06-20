package router

import (
	"iv-go/apis"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由设置
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/welcome", apis.Welcome)
	return router
}
