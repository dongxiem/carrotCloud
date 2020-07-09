package controllers

import (
	"carrotCloud/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserLogin：用户登录，进行验证
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	// 参数接收正确, 进行登录操作
	if err := c.ShouldBindJSON(&service); err != nil {
		res := service.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		// 验证错误
		c.JSON(http.StatusUnprocessableEntity, ErrorResponse(err))
	}

}
