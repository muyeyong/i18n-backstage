/*
 * @Author: xuyong
 * @Date: 2024-08-08 20:05:39
 * @LastEditors: xuyong
 */
package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&User{}, &Video{}, &Language{}, &Translation{})
}
