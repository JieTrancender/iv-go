package apis

import (
	"math"
	"net/http"
	"strconv"

	"iv-go/models"
	"iv-go/pkg"

	"github.com/gin-gonic/gin"
)

// UserIndex 用户列表
func UserIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	data := models.GetUsers(page, size)

	meta := make(map[string]interface{})
	total, _ := models.GetUserTotal()
	meta["total"] = total
	meta["current_page"] = page
	meta["per_page"] = size
	meta["last_page"] = math.Ceil(float64(total/size)) + 1

	pkg.RespJSON(c, http.StatusOK, "success", gin.H{
		"meta": meta,
		"list": data,
	})
}
