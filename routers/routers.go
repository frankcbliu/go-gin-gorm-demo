package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-demo/utils"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/register", UserRegister) // 注册
		user.GET("/login", UserLogin)       // 登录
	}
	//user.Use(utils.CookieCheck())
	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "小明", "value": "pong"})
	})

	r.GET("/pong", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/login", func(c *gin.Context) {
		// Set cookie {"label": "ok" }, maxAge 30 seconds.
		utils.Success(c, gin.H{"data": "Your home page"}, "success")
	})

	r.GET("/home", utils.CookieCheck(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Your home page"})
	})
	return r
}
