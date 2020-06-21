package pkg

import "github.com/gin-gonic/gin"

// RespJSON for resp json data
func RespJSON(c *gin.Context, code int, message string, content interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    content,
	})
}
