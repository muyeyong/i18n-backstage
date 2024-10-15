/*
 * @Author: xuyong
 * @Date: 2024-10-15 19:19:49
 * @LastEditors: xuyong
 */
package model

import (
	"gorm.io/gorm"
)

type Translation struct {
	gorm.Model	
	SourceLanguageID uint
	TargetLanguageID uint
	SourceText string
	TTranslatedText string
}