package myutils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// 失败返回函数
func ReturnFail(ctx *app.RequestContext, scode int, msg string) {
	ctx.JSON(scode, utils.H{
		"status":  scode,
		"message": msg,
	})
}

// 成功返回函数
func ReturnSuccess(ctx *app.RequestContext, scode int, msg string) {
	ctx.JSON(scode, utils.H{
		"status":  scode,
		"message": msg,
	})
}
