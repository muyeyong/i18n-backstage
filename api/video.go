/*
 * @Author: xuyong
 * @Date: 2024-08-12 17:01:56
 * @LastEditors: xuyong
 */
package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 创建视频
func CreateVideo(c *gin.Context) {
	// 需要将数据存入到数据库中
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// ShowVideo 视频详情 
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListVideo 视频列表
func ListVideo (c *gin.Context) {
	service := service.ListVideoService{}
	res := service.List()
	c.JSON(200, res)
}

// UpdateVideo 更新视频
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}