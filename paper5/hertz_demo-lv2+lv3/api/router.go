package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go_papers/paper5/hertz_demo-lv1/middleware"
)

// 统一路由
func InitRouter() {
	h := server.New(server.WithHostPorts(":8080")) //拉取引擎，自定义端口

	// 添加日志中间件，终端打印请求信息
	h.Use(middleware.PrintLogger())

	h.POST("/addstu", Addapi)
	h.POST("/profile", Profileapi)
	h.GET("/search", Searchapi)
	// 监听端口
	h.Spin()
}
