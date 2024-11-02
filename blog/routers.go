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
	r.POST("api/v1/category/create", middleware.AuthMiddleware(), controller.CreateCategory)
	r.DELETE("api/v1/category/delete", middleware.AuthMiddleware(), controller.DeleteCategory)
	r.GET("api/v1/category/categories", middleware.AuthMiddleware(), controller.GetCategories)

	//post相关的接口
	r.POST("api/v1/post/create", middleware.AuthMiddleware(), controller.CreatePost)
	return r
}
