package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	//name, _ := c.FormFile("file")
	file, fileHeader, err := c.Request.FormFile("File")
	//fmt.Printf("file:%v\nfileHeader:%v\nerr;%v\n", file, fileHeader, err)
	code := errmsg.ERROR
	//c.Request.
	if err != nil {
		c.JSON(200, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
