/*
 * @Author: xuyong
 * @Date: 2024-08-13 20:09:39
 * @LastEditors: xuyong
 */
package service

import (
	"singo/model"
	"singo/serializer"
)

type CreateVideoService struct {
	Title string `form: "title" json:"title" binding:"required,min=2,max=40"`
	Info string `form: "info" json: "info" binding:"required, max=300`
}

func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video {
		Title: service.Title,
		Info: service.Info,
	}
err := model.DB.Create(&video).Error
if err != nil {
	return serializer.Response {
		Code: 50001,
		Msg: "视频创建失败",
	}
}
	return serializer.BindVideoResponse(video)
}