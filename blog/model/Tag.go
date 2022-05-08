package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Tag struct {
	Name  string `gorm:"type:varchar(20)" json:"name"`
	Color string `gorm:"type:varchar(10)" json:"color"`
	gorm.Model
}

// 检查标签是否存在
func CheckTagIsExist(name string) int {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return errmsg.ERROR_TAG_EXIST
	}
	return errmsg.SUCCESS
}

// 添加标签
func CreateTag(data *Tag) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除标签
func DeleteTag(id int) int {
	var tag Tag
	err = db.Where("id = ?", id).Delete(&tag).Error
	// 删除映射关系,后续文章获取所有的标签的话是通过这个进行获取的
	db.Model(&ArtToTag{}).Where("tid = ?", id).Delete(&ArtToTag{})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑标签,编辑标签的名字和颜色
func EditTag(id int, data *Tag) int {
	var tag Tag
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["color"] = data.Color
	err = db.Model(&tag).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询标签列表
func SelectAllTages() ([]Tag, int) {
	var tags []Tag
	err = db.Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return tags, errmsg.SUCCESS
}

/*
 * TODO：
 * 下面的这些函数是辅助文章编写的时候使用
 * 1. 为文章添加标签
 * 2. 编辑文章的时候作者可能需要删除部分标签或者重命名（修改）标签
 * 3. 查询文章的同时需要查到该文章的所有标签
 */
