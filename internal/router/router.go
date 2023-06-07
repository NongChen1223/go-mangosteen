package router

import (
	"github.com/gin-gonic/gin"
	"mangosteen/internal/controller"
)

/*
go语言没有构造函数的概念，一般会使用NewXXX来初始化相关类。
*/
func New() *gin.Engine {
	r := gin.Default()

	// Ping test 以路由为单位进行分组分发请求
	r.GET("/ping", controller.Ping)
	return r
}
