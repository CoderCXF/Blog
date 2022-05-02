package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	// Belong to（many to one）关系
	Cid      int      `gorm:"type:int;not null;" json:"cid"`
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	// 分类id
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
	gorm.Model
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 查询指定分类下的所有文章
func GetCatArt(pageSize int, pageNum int, cid int) ([]Article, int, int64) {
	var articles []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).
		Find(&articles).Count(&total).Error
	if err != nil {
		return articles, errmsg.ERROR_CAT_NOT_EXIST, 0
	}
	return articles, errmsg.SUCCESS, total
}

// 查询单篇文章详情
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// 查询文章列表 返回Article类型的切片
func GetArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var total int64
	// 分页
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articles, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// TODO:查询该分类下的所有文章

// 删除分类(软删除)
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
