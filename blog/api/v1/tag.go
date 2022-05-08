package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加标签
func AddTag(c *gin.Context) {
	var tag model.Tag
	_ = c.ShouldBindJSON(&tag)
	name := tag.Name
	code = model.CheckTagIsExist(name)
	if code == errmsg.SUCCESS {
		code = model.CreateTag(&tag)
	}
	c.JSON(200, gin.H{
		"status":  code,
		"data":    tag,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除标签，根据param id参数
// 不需要判断tag是否存在
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteTag(id)
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑标签
func EditTag(c *gin.Context) {
	var data model.Tag
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.EditTag(id, &data)
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询标签列表
func GetAllTags(c *gin.Context) {
	data, code := model.SelectAllTages()
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
