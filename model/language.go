/*
 * @Author: xuyong
 * @Date: 2024-10-15 19:23:04
 * @LastEditors: xuyong
 */
package model

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	LanguageName  string 
}