/*
 * @Author: xuyong
 * @Date: 2024-10-12 14:59:46
 * @LastEditors: xuyong
 */
package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowVideoService struct {}

func (service *ShowVideoService)Show(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "视频不存在",
			Error: err.Error(),
		}
	}
	return serializer.BindVideoResponse(video)
}