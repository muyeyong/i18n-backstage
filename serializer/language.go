/*
 * @Author: xuyong
 * @Date: 2024-10-15 19:55:44
 * @LastEditors: xuyong
 */
package serializer

import "singo/model"

type Language struct {
	ID uint `json:"id"`
	LanguageName string `json:"language_name"`
	CreatedAt int64 `json:"created_at"`
}

func BuildLanguage (lan model.Language) Language {
	return Language {
		ID: lan.ID,
		LanguageName: lan.LanguageName,
		CreatedAt: lan.CreatedAt.Unix(),
	}
}
func BuildLanguageResponse(lan model.Language) Response {
	return Response {
		Data: BuildLanguage(lan),
	}
}