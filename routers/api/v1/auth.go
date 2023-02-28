package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok,_ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		if models.CheckAuth(username, password) {
			token, err := util.GenerateToken(username, password)
			if err == nil {
				data["token"] = token
				code = e.SUCCESS
			}else {
				code = e.ERROR_AUTH_TOKEN
			}
		}else {
			code = e.ERROR_AUTH
		}
	}else {
		//fmt.Printf("auth error %+v", valid.Errors)
		logging.Error(valid.Errors)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}