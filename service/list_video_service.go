/*
 * @Author: xuyong
 * @Date: 2024-10-12 16:00:16
 * @LastEditors: xuyong
 */
package service

import (
	"singo/model"
	"singo/serializer"
)

type ListVideoService struct{}

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response {
			Code: 50001,
			Msg:  "视频列表获取失败",
			Error: err.Error(),
		}
	} else {
		return serializer.BindVideosResponse(videos)
	}
}