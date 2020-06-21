package apis

import (
	"net/http"

	"iv-go/models"

	"github.com/gin-gonic/gin"
)

// UserIndex 用户列表
func UserIndex(c *gin.Context) {
	var user models.User
	result := models.DB.Take(&user).Value

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"content": result,
		},
	})
}
