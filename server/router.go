/*
 * @Author: xuyong
 * @Date: 2024-08-08 20:05:39
 * @LastEditors: xuyong
 */
package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// createvideoservice
// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	// 保持登录状态
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	// 跨域
	r.Use(middleware.Cors())
	// 验证登录
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		v1.POST("videos", api.CreateVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)

		v1.POST("language/:name", api.CreateLanguage)
	}
	return r
}
