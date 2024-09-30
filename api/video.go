/*
 * @Author: xuyong
 * @Date: 2024-08-12 17:01:56
 * @LastEditors: xuyong
 */
package api

import (
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

// CreateVideo 创建视频
func CreateVideo(c *gin.Context) {
	// 需要将数据存入到数据库中
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:"创建视频成功",
	})
}
