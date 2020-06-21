package router

import (
	"fmt"
	"html/template"
	"iv-go/apis"
	"net/http"
	"time"

	"iv-go/pkg"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由设置
func SetupRouter() *gin.Engine {
	router := gin.New()

	// web页面路由
	pageGroup(router)

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

func pageGroup(r *gin.Engine) *gin.Engine {
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	r.LoadHTMLGlob("views/**/*")

	page := r.Group("/page", pkg.EnableCookieSession())
	{
		page.StaticFS("/assets", http.Dir("assets"))
		page.GET("/raw", func(c *gin.Context) {
			c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
				"now": time.Date(2020, 06, 21, 0, 0, 0, 0, time.UTC),
			})
		})
	}

	return r
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
