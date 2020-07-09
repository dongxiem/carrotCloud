package service

import (
	"carrotCloud/pkg/serializer"
	"github.com/gin-gonic/gin"
)

// UserLoginService： 用户登录管理服务
// 绑定模型
type UserLoginService struct {
	// 用户名长度最小为5位，最大为30位
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	// 密码长度最小8位，最大40位
	Password    string `form:"Password" json:"password" bingding:"required,min=8,max=40"`
	CaptchaCode string `form:"captchaCode" json:"captcha_code"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	return serializer.Response{
		Code: 0,
		Msg:  "OK",
	}
}
