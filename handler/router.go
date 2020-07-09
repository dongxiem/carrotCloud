package handler

import (
	"carrotCloud/handler/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	rout := r.Group("/Api/carrot")
	{
		rout.GET("Ping", controllers.Ping)
	}
	return r
}
