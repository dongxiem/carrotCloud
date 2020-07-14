package route

import (
	hdl "carrotCloud/handler"
	midware "carrotCloud/middleWare"

	"github.com/gin-gonic/gin"
)

// Router ： 路由配置
func Router() *gin.Engine {
	// gin framework
	router := gin.Default()

	// 静态资源处理
	router.Static("/static/", "./static")

	// 不需验证的接口
	router.GET("/user/signup", hdl.SignUpHandler)
	router.GET("/user/signin", hdl.SignInHandler)
	router.POST("/user/signup", hdl.DoSignUpHandler)
	router.POST("/user/signin", hdl.DoSignInHandler)
	router.GET("/user/exists", hdl.UserExistsHandler)

	// 加入auth认证中间件
	router.Use(midware.Authorize())

	// 用户相关接口
	router.POST("/user/info", hdl.UserInfoHandler)

	return router
}
