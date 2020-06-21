package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"iv-go/router"

	"iv-go/models"

	"github.com/gin-gonic/gin"
)

// Login for login binding and validation
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func setupRouter() *gin.Engine {
	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(gin.LoggerWithFormatter(loggerFormatter))

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "iv")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "iv")

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": message,
			"nick":    nick,
		})
	})

	// /map_post?ids[a]=123 -d 'names[pr]=iv'
	router.POST("/map_post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v, names: %v", ids, names)
		c.String(http.StatusOK, fmt.Sprintf("ids: %v, names: %v", ids, names))
	})

	router.MaxMultipartMemory = 8 << 20 // 8MiB, default is 32MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename, file.Header)

		err := c.SaveUploadedFile(file, "./file/"+file.Filename)
		if nil != err {
			log.Println("err:", err)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
	})

	router.POST("/uploadMulti", uploadMulti)

	v1 := router.Group("/v1")
	{
		v1.POST("/login", defaultHandler)
		v1.POST("/submit", defaultHandler)
		v1.POST("/read", defaultHandler)

		v1.POST("/loginJSON", loginJSONHandler)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", defaultHandlerV2)
		v2.POST("/submit", defaultHandlerV2)
		v2.POST("/read", defaultHandlerV2)
	}

	router.Static("/static", "./static")

	return router
}

func init() {
	models.Setup()
}

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("log/iv.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := router.SetupRouter()
	// router := setupRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func uploadMulti(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "./file/"+file.Filename)
	}

	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
}

func defaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"method": c.Request.Method,
			"url":    c.Request.URL.Path,
		},
	})
}

func defaultHandlerV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"method": c.Request.Method,
			"url":    c.Request.URL.Path,
		},
	})
}

func loggerFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n", param.ClientIP,
		param.TimeStamp.Format(time.RFC1123), param.Method, param.Path, param.Request.Proto,
		param.StatusCode, param.Latency, param.Request.UserAgent(), param.ErrorMessage)
}

func loginJSONHandler(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	if json.User != "root" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "unauthorized",
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"username": json.User,
		},
	})
}
