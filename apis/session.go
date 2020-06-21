package apis

import (
	"iv-go/models"
	"iv-go/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {
	var user models.User
	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")

	if hasSession := pkg.HasSession(c); hasSession == true {
		pkg.RespJSON(c, http.StatusOK, "fail", "user has logind")
		return
	}

	if existUser := models.UserDetailByName(user.Name); existUser.ID != 0 {
		pkg.RespJSON(c, http.StatusOK, "fail", "username has existed")
		return
	}

	if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
		pkg.RespJSON(c, http.StatusOK, "fail", "inconsistent password")
		return
	}

	if pwd, err := pkg.Encrypt(c.Request.FormValue("password")); err == nil {
		user.Password = pwd
	}

	models.AddUser(&user)

	pkg.SaveAuthSession(c, user.ID)

	pkg.RespJSON(c, http.StatusOK, "success", "register success")
}

// Login 登录
func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")

	if hasSession := pkg.HasSession(c); hasSession == true {
		pkg.RespJSON(c, http.StatusOK, "fail", "user has logind")
		return
	}

	user := models.UserDetailByName(name)

	if err := pkg.Compare(user.Password, password); err != nil {
		pkg.RespJSON(c, http.StatusOK, "fail", "name or password is wrong")
		return
	}

	pkg.SaveAuthSession(c, user.ID)

	pkg.RespJSON(c, http.StatusOK, "success", "login success")
}

// Logout 登出
func Logout(c *gin.Context) {
	if hasSession := pkg.HasSession(c); hasSession == false {
		pkg.RespJSON(c, http.StatusOK, "fail", "user is not being login status")
		return
	}

	pkg.ClearAuthSession(c)

	pkg.RespJSON(c, http.StatusOK, "success", "logout success")
}

// Me 用户自己信息
func Me(c *gin.Context) {
	curUser := c.MustGet("userId").(uint)
	pkg.RespJSON(c, http.StatusOK, "success", curUser)
}
