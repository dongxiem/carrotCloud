package rout

import (
	hdl "carrotCloud/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin framework
	router := gin.Default()

	// 用户相关接口
	router.POST("/user/info", hdl.UserInfoHandler)

	return router
}
