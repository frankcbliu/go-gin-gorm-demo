package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-demo/utils"
)

func UserRegister(c *gin.Context) {

}

func UserLogin(c *gin.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")
	utils.SetCookie(c, userName)
	println(userName, password)
	utils.Success(c, gin.H{}, "login success")
}
