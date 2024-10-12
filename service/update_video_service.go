/*
 * @Author: xuyong
 * @Date: 2024-10-12 17:20:16
 * @LastEditors: xuyong
 */
package service

import (
	"singo/model"
	"singo/serializer"
)

type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=40"`
	Info string `form:"info" json:"info" binding:"required,max=300"` 
}

func (service *UpdateVideoService)Update(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.Find(&video, id).Error
	if err != nil {
		return serializer.Response {
			Code: 50001,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response {
			Code: 50001,
			Msg: "视频更新失败",
			Error: err.Error(),
		}
	}
	return serializer.BindVideoResponse(video)
}