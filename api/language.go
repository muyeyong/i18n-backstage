package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// 新增语言
func CreateLanguage(c *gin.Context) {
	service := service.CreateLanguageService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 修改语言

// 删除语言

// 更新语言

// 语言列表