package apis

import (
	"iv-go/app"
	"iv-go/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome 欢迎
func Welcome(c *gin.Context) {
	pkg.RespJSON(c, http.StatusOK, "success", app.AppEnv.IV.Welcome)
}
