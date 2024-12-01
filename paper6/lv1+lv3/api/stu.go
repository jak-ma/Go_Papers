package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go_papers/paper6/lv1+lv3/dao"
	"go_papers/paper6/lv1+lv3/model"
	"go_papers/paper6/lv1+lv3/myutils"
)

// 添加学生接口
func Addapi(c context.Context, ctx *app.RequestContext) {
	sex := ctx.PostForm("sex")
	name := ctx.PostForm("name")
	comefrom := ctx.PostForm("comefrom")

	if !dao.AddStudent(sex, name, comefrom) {
		myutils.ReturnFail(ctx, 400, "add failing")
	} else {
		myutils.ReturnSuccess(ctx, 200, "add success")
	}
}

// 删除学生接口
func Deleteapi(c context.Context, ctx *app.RequestContext) {
	name := ctx.PostForm("name")
	if !dao.Delete(name) {
		myutils.ReturnFail(ctx, 400, "delete fail")
	} else {
		myutils.ReturnSuccess(ctx, 200, "delete success")
	}
}

// 修改学生信息
func Profileapi(c context.Context, ctx *app.RequestContext) {
	var S model.Student
	S.Name = ctx.PostForm("name")
	S.Sex = ctx.PostForm("sex")
	S.Comefrom = ctx.PostForm("comefrom")
	if !dao.Profile(S) {
		myutils.ReturnFail(ctx, 400, "fix fail")
		return
	} else {
		myutils.ReturnSuccess(ctx, 200, "fix success")
	}
}

// 指定删除某一信息
func DeleteSingleapi(c context.Context, ctx *app.RequestContext) {
	args := ctx.PostForm("args")
	name := ctx.PostForm("name")
	if !dao.DeleteSingle(args, name) {
		myutils.ReturnFail(ctx, 400, "delete fail or time write error")
		return
	} else {
		myutils.ReturnSuccess(ctx, 200, "delete success")
	}
}
