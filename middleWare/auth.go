package middleWare

import (
	"carrotCloud/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authorize : 权鉴功能-HTTP拦截器
func Authorize() gin.HandlerFunc {

	return func(c *gin.Context) {
		username := c.Request.FormValue("username") // 用户名
		token := c.Request.FormValue("token")       // 访问令牌
		fmt.Println(username)
		fmt.Println(len(token))
		// 进行长度和Token的验证
		if len(username) < 3 || !handler.IsTokenValid(token) {
			// 验证不通过，不再调用后续的函数进行处理了
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			// return可以进行省略，因为Abort()就可以让后面的Handler函数不再执行了
			return
		}
		c.Next()
	}
}
