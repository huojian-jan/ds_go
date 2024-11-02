package main

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("api/auth/login", controller.LogIn)
	//获取用信息，应该需要授权，这里先跳过
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
