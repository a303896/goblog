package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1
	if param := c.Query("state"); param != "" {
		state = com.StrTo(param).MustInt()
		maps["state"] = state
	}
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称不能超过100个字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人不能超过100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只能为0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			tag := models.Tag{
				Name:       name,
				CreatedBy:  createdBy,
				State:      state,
			}
			code = e.SUCCESS
			models.AddTag(tag)
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	name := c.Query("name")
	id := com.StrTo(c.Param("id")).MustInt()
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称不能超过100个字符")
	valid.Required(modifiedBy, "modified_by").Message("创建人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("创建人不能超过100个字符")
	valid.Required(id, "id").Message("标签ID不能为空")
	fmt.Printf("name=%s,modifiedBy=%s\n", name, modifiedBy)
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			data := make(map[string]interface{})
			if name != "" {
				data["name"] = name
			}
			if modifiedBy != "" {
				data["modified_by"] = modifiedBy
			}
			models.UpdateTagById(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}else {
		fmt.Printf("errors:%+v\n",valid.Errors)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("缺少标签ID")

	if !valid.HasErrors() {
		models.DeleteTagById(id)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
		"data": make(map[string]string),
	})
}
