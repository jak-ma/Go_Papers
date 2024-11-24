package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go_papers/paper5/hertz_demo-lv2+lv3/dao"
	"go_papers/paper5/hertz_demo-lv2+lv3/myutils"
)

// 添加学生接口
func Addapi(c context.Context, ctx *app.RequestContext) {
	dao.LoadDataFromFile()
	//dao.LoadDataFromJson()
	sex := ctx.PostForm("sex")
	name := ctx.PostForm("name")
	id := ctx.PostForm("id")
	comefrom := ctx.PostForm("comefrom")
	if dao.SearchId(id) {
		myutils.ReturnFail(ctx, 400, "current id is already existed")
		return
	}
	dao.AddStudent(sex, name, id, comefrom)
	myutils.ReturnSuccess(ctx, 200, "add success")
}

// 修改信息接口
func Profileapi(c context.Context, ctx *app.RequestContext) {
	dao.LoadDataFromFile()
	//dao.LoadDataFromJson()
	id := ctx.PostForm("id")
	newcomefrom := ctx.PostForm("comefrom")
	if !dao.SearchId(id) {
		myutils.ReturnFail(ctx, 400, "input id isn't existed")
		return
	}
	dao.ProfileStuInfro(id, newcomefrom)
	myutils.ReturnSuccess(ctx, 200, "stu information profile successfully")
}

// 查询信息接口
func Searchapi(c context.Context, ctx *app.RequestContext) {
	dao.LoadDataFromFile()
	//dao.LoadDataFromJson()
	id := ctx.Query("id")
	if !dao.SearchId(id) {
		myutils.ReturnFail(ctx, 404, "the stu of current input id isn't exist")
		return
	}
	stu := dao.Data[id]
	ctx.JSON(200, utils.H{
		"msg":         "success",
		"stuname":     stu.Name,
		"stuid":       stu.Id,
		"stusex":      stu.Sex,
		"stucomefrom": stu.Comefrom,
		"optime":      stu.Optime,
	})
}
