package main

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	//user相关的接口
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("api/v1/auth/register", controller.Register)
	r.POST("api/v1/auth/login", controller.LogIn)
	//获取用信息，应该需要授权，这里先跳过
	r.GET("api/v1/auth/info", middleware.AuthMiddleware(), controller.Info)

	//category相关的接口
	r.POST("api/v1/category/create", controller.CreateCategory)
	r.GET("api/v1/category/get", middleware.AuthMiddleware(), controller.GetCategoryById)

	return r
}
