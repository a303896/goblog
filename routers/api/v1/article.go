package v1

import (
	"blog/models"
	"blog/pkg/e"
	"fmt"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticles(c *gin.Context)  {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	page := com.StrTo(c.Query("page")).MustInt()
	pageSize := com.StrTo(c.Query("limit")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	name := c.Query("name")

	if tagId > 0 {
		maps["tag_id"] = tagId
	}

	if name != "" {
		maps["name"] = name
	}
	data["list"] = models.GetArticles(page, pageSize, maps)
	data["total"] = models.GetArticleTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

func GetArticle(c *gin.Context)  {
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id,"id")
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleById(id) {
			data = models.GetArticleById(id)
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		fmt.Printf("errors:%+v\n",valid.Errors)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

func EditArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")
	state := com.StrTo(c.PostForm("state")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id")
	valid.Required(title, "title")
	valid.Required(desc, "desc")
	valid.Required(content, "content")
	valid.Required(modifiedBy, "modified_by")
	valid.Required(state, "state")
	valid.Min(tagId, 1, "tag_id")
	valid.Range(state, 0, 1, "state")

	code := e.INVALID_PARAMS
	var article models.Article
	if !valid.HasErrors() {
		code = e.SUCCESS
		article = models.GetArticleById(id)
		if article != (models.Article{}) {
			article.Title = title
			article.Desc = desc
			article.ModifiedBy = modifiedBy
			article.Content = content
			article.State = state
			if tagId > 0 && models.ExistTagById(tagId) {
				article.TagID = tagId
			}
			models.UpdateArticleById(id, article)
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		fmt.Printf("errors:%+v\n",valid.Errors)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": article,
	})
}

func AddArticle(c *gin.Context)  {
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	state := com.StrTo(c.PostForm("state")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title")
	valid.Required(desc, "desc")
	valid.Required(content, "content")
	valid.Required(createdBy, "created_by")
	valid.Required(state, "state")
	valid.Range(state, 0, 1, "state")

	code := e.INVALID_PARAMS
	article := models.Article{}
	if !valid.HasErrors() {
		code = e.SUCCESS
		article = models.Article{
			Title: title,
			Desc: desc,
			Content: content,
			CreatedBy: createdBy,
			State: state,
		}
		if tagId > 0 && models.ExistTagById(tagId) {
			article.TagID = tagId
		}
		models.CreateArticle(&article)
	}else {
		fmt.Printf("errors:%+v\n",valid.Errors)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": article,
	})
}

func DeleteArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id")
	valid.Min(id, 1, "id")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleById(id) {
			models.DeleteArticleById(id)
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}


