package models

import (
	"fmt"
	"gorm.io/plugin/soft_delete"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_on;type:int"`
	State int `json:"state"`
}

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//	return nil
//}

func ExistArticleById(id int) bool {
	var result Article
	db.Select("id").Where("id = ?", id).First(&result)
	if result.ID > 0 {
		return true
	}
	return false
}

func CreateArticle(article *Article) bool {
	res := db.Create(article)
	if res.Error == nil {
		return true
	}
	fmt.Printf("添加文章发生错误：%s", res.Error)
	return false
}

func UpdateArticleById(id int, data interface{}) bool {
	res := db.Model(new(Article)).Where("id = ?", id).Updates(data)
	if res.Error == nil {
		return true
	}
	fmt.Printf("更新文章发生错误：%s", res.Error)
	return false
}

func GetArticleTotal(where interface{}) (value int) {
	db.Model(new(Article)).Where(where).Count(&value)
	return
}

func GetArticleById(id int) (article Article) {
	db.Model(new(Article)).Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(page int, pageSize int, where interface{}) (list []Article) {
	db.Preloads("Tag").Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&list)
	return
}

func DeleteArticleById(id int) bool {
	db.Where("id = ?", id).Delete(new(Article))
	return true
}

func ClearAllArticle() bool {
	db.Unscoped().Where("delete_on > ?", 0).Delete(new(Article))
	return true
}