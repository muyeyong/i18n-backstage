/*
 * @Author: xuyong
 * @Date: 2024-10-14 16:11:05
 * @LastEditors: xuyong
 */
package service

import (
	"singo/model"
	"singo/serializer"
)

type DeleteVideoService struct {
}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.Find(&video, id).Error
	if err != nil {
		return serializer.Response {
			Code: 50001,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	} else {
		err = model.DB.Delete(&video).Error
		if err != nil {
			return serializer.Response {
				Code: 50001,
				Msg: "视频删除失败",
				Error: err.Error(),
			}
		} else {
			return serializer.Response {
				Code: 0,
				Msg: "视频删除成功",
			}
		}
	}
}