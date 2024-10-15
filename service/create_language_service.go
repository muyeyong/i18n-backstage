/*
 * @Author: xuyong
 * @Date: 2024-10-15 19:36:28
 * @LastEditors: xuyong
 */
package service

import (
	"singo/model"
	"singo/serializer"
)

type CreateLanguageService struct {
	LanguageName string `form:"language_name" json:"language_name" binding:"required"`
}

func (service *CreateLanguageService) Create() serializer.Response {
	language := model.Language{
		LanguageName: service.LanguageName,
	}
	err := model.DB.Create(&language).Error
	if err != nil {
		return serializer.Response {
			Code: 50001,
			Msg: "创建语言失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildLanguageResponse(language)
}