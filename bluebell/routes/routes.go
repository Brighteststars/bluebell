package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册路由

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controllers.SignUpHandler)
	// 登录
	r.POST("/login", controllers.LoginHandler)

	r.GET("/luffy", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "luffy")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
