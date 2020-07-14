package handler

import (
	"carrotCloud/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	dblayer "carrotCloud/db"
)

const (
	// 加盐
	pwdSalt = "*#890"
)

// SignUpHandler : 处理用户注册请求
func SignUpHandler(c *gin.Context) {
	// 进行重定向
	c.Redirect(http.StatusFound, "http://"+c.Request.Host+"/static/view/signup.html")

}

// DoSignUpHandler : 处理用户注册请求
func DoSignUpHandler(c *gin.Context) {
	// 获取用户名字和密码
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	// 参数长度判断，进行校验
	if len(username) < 3 || len(passwd) < 5 {
		// 直接填入JSON数据并且返回
		c.JSON(http.StatusOK,
			gin.H{
				"msg": "Invalid parameter",
			})
		return
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	// 将用户信息注册到用户表中
	suc := dblayer.UserSignUp(username, encPasswd)
	if suc {
		c.JSON(http.StatusOK,
			gin.H{
				"code":    0,
				"msg":     "注册成功",
				"data":    nil,
				"forward": "/user/signin",
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"code": 0,
				"msg":  "注册失败",
				"data": nil,
			})
	}
}

// SignInHandler : 处理用户登录请求
func SignInHandler(c *gin.Context) {
	// 进行重定向
	c.Redirect(http.StatusFound, "http://"+c.Request.Host+"/static/view/signin.html")
}

// DoSignInHandler : 登录接口
func DoSignInHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	encPasswd := util.Sha1([]byte(password + pwdSalt))

	// 1.校验用户名及密码
	pwdChecked := dblayer.UserSignIn(username, encPasswd)
	if !pwdChecked {
		// 写入JSON，验证失败
		c.JSON(http.StatusOK,
			gin.H{
				"code": 0,
				"msg":  "密码校验失败",
				"data": nil,
			})
		return
	}
	// 2.生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		// 写入JSON，登录失败信息
		c.JSON(http.StatusOK,
			gin.H{
				"code": 0,
				"msg":  "登录失败",
				"data": nil,
			})
		return
	}
	// 3.登录成功之后重定向到首页
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			// 这里的c.Request.Host 挺好用
			Location: "http://" + c.Request.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	// 写入数据
	c.Data(http.StatusOK, "octet-stream", resp.JSONBytes())
}

// UserInfoHandler ： 查询用户信息
func UserInfoHandler(c *gin.Context) {
	// 1.解析参数请求
	username := c.Request.FormValue("username")

	// 2.查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		c.JSON(http.StatusForbidden,
			gin.H{})
		return
	}

	// 3.组装并响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	// 写入数据
	c.Data(http.StatusOK, "octet-stream", resp.JSONBytes())
}

// GenToken ： 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// IsTokenValid : 判断token是否有效
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断token的时效性，是否过期
	// TODO: 从数据库表tbl_user_token查询username对应的token信息
	// TODO: 对比两个token是否一致
	return true
}

// UserExistsHandler : 查询用户是否存在
func UserExistsHandler(c *gin.Context) {

	// 1.解析请求参数
	username := c.Request.FormValue("username")

	// 2.查询用户信息
	exists, err := dblayer.UserExist(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"msg": "server error",
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"msg":    "ok",
				"exists": exists,
			})
	}
}
