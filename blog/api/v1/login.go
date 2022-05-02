package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

/*登录页面设计，使用jwt进行身份认证*/
// 1. 登录验证用户名和密码
// 2. 生成token返回客户端
func Login(c *gin.Context) {
	var data model.User
	var token string
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.Username, data.Password)
	// 验证成功可以登录
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
