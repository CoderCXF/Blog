package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	val "blog/utils/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 控制器

var code int

// 查询用户是否存在

func UserExist(c *gin.Context) {

}

// 添加用户:POST请求

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 注册校验
	msg, code := val.Validate(data)
	if code != errmsg.SUCCESS {
		c.JSON(200, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		// 可以添加数据
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户（无意义）

// 查询用户列表

func GetUsers(c *gin.Context) {
	// c.Query表示获取前端的请求参数，
	// 比如带有参数的请求/path?id=1234&name=Manu&value=111,
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// 当参数默认为0的话
	//if pageSize == 0 {
	//	pageSize = -1
	//}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
// FIXME: 该逻辑如果不修改用户名的话，则其他信息没办法修改，BUG

func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
