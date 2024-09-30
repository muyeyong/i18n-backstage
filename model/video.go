/*
 * @Author: xuyong
 * @Date: 2024-08-12 17:09:29
 * @LastEditors: xuyong
 */
package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title string
	Info string
}