package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go_papers/paper5/hertz_demo-lv1/middleware"
)

func main() {
	h := server.New(server.WithHostPorts(":8080")) //自定义端口

	// 添加日志中间件
	h.Use(middleware.PrintLogger())

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.GET("/echo", func(ctx context.Context, c *app.RequestContext) {
		name := c.Query("message")
		c.JSON(consts.StatusOK, utils.H{
			"name": name,
		})
	})
	// 监听端口
	h.Spin()
}
