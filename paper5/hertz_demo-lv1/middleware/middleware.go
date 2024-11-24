package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

// 打印访问状态
func PrintLogger() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		method := ctx.Method()
		path := ctx.Path()
		remoteAddr := ctx.RemoteAddr()
		ctx.Next(c)
		status := ctx.Response.StatusCode()
		fmt.Printf("%s %d %s %s\n", method, status, path, remoteAddr)

	}
}
