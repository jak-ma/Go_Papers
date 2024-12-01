package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go_papers/paper6/lv1+lv3/myutils"
	"go_papers/paper6/lv2+lv3/dao"
)

// register
func Registerapi(c context.Context, ctx *app.RequestContext) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" || password == "" {
		myutils.ReturnFail(ctx, 400, "username or password doesn't be empty")
		return
	}
	if dao.SearchUser(username) {
		myutils.ReturnFail(ctx, 400, "username already exist or mysql read fail")
		return
	}
	if !dao.AddUser(username, password) {
		myutils.ReturnFail(ctx, 500, "mysql write fail")
		return
	} else {
		myutils.ReturnSuccess(ctx, 200, "register success")
	}
}

// login
func Loginapi(c context.Context, ctx *app.RequestContext) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" || password == "" {
		myutils.ReturnFail(ctx, 400, "username or password doesn't be empty")
		return
	}
	if !dao.SearchUser(username) {
		myutils.ReturnFail(ctx, 400, "username is doesn't exist")
		return
	}
	if dao.SearchPwdFromUsername(username, password) {
		myutils.ReturnSuccess(ctx, 200, "login success")
		return
	} else {
		myutils.ReturnFail(ctx, 400, "login fail,password error or mysql read fail")
	}
}

// set security_question/security_answer
func SetSecurityQuapi(c context.Context, ctx *app.RequestContext) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	security_question := ctx.PostForm("security_question")
	security_answer := ctx.PostForm("security_answer")
	// 省略掉输入为nil的情况判断
	if dao.SetSecurityQu(username, password, security_question, security_answer) {
		myutils.ReturnSuccess(ctx, 200, "set security .. success")
		return
	} else {
		myutils.ReturnFail(ctx, 400, "set fail")
		return
	}
}

// update password
func Resetapi(c context.Context, ctx *app.RequestContext) {
	username := ctx.PostForm("username")
	newpassword := ctx.PostForm("newpassword")
	security_answer := ctx.PostForm("security_answer")
	if dao.ResetPassword(username, newpassword, security_answer) {
		myutils.ReturnSuccess(ctx, 200, "reset success")
		return
	} else {
		myutils.ReturnFail(ctx, 400, "reset fail")
		return
	}
}
