package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//添加Get请求
	router.GET("/", func(cxt *gin.Context) {
		cxt.String(http.StatusOK, "hello gin")
	})
	return router
}
