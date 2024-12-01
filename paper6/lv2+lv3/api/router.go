package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go_papers/paper5/hertz_demo-lv1/middleware"
)

func InitRouter() {
	h := server.New(server.WithHostPorts(":8080"))
	h.Use(middleware.PrintLogger())

	h.POST("/register", Registerapi)

	h.POST("/login", Loginapi)

	h.POST("/set", SetSecurityQuapi)

	h.POST("/resetpwd", Resetapi)

	h.Spin()
}
