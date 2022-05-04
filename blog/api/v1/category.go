package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 控制器

// 添加分类:POST请求

func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		// 可以添加数据
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCateInfo(id)
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章
// TODO:

// 查询分类列表：GET

func GetCates(c *gin.Context) {
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
	data, total := model.GetCate(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类名：PUT

func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类：DELETE

func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
