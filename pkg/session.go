package pkg

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"iv-go/models"
)

// KEY for gin session
const KEY = "keyboard"

// EnableCookieSession 使用cookie存储session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("iv-go", store)
}

// AuthSessionMiddle middle for auth session
func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"message": "unauthorized",
				"content": gin.H{},
			})
			c.Abort()
			return
		}

		c.Set("userId", sessionValue.(uint))

		c.Next()
		return
	}
}

// SaveAuthSession 存储session信息
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

// ClearAuthSession 清理session信息
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// HasSession session是否存在
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}

	return true
}

// GetSessionUserID 从session获取用户id
func GetSessionUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}

	return sessionValue.(uint)
}

// GetUserSession 获取用户session
func GetUserSession(c *gin.Context) map[string]interface{} {
	hasSession := HasSession(c)
	userName := ""
	if hasSession {
		userID := GetSessionUserID(c)
		userName = models.UserDetail(userID).Name
	}

	data := make(map[string]interface{})
	data["hasSession"] = hasSession
	data["userName"] = userName
	return data
}