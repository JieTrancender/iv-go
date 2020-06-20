package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome 欢迎
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"content": "hello welcome",
		},
	})
}
