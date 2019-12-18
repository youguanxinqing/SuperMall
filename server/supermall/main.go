package main

import (
	"github.com/kataras/iris/v12"

	"supermall/middleware"
	"supermall/views"
)

func init() {
}

func main() {
	app := iris.Default()

	// 中间件设置
	app.Use(middleware.AccessControlAllowOrigin)

	/* 注册路由 */

	// 静态资源
	app.Get("/static/{filepath:path}", views.StaticResource)

	// 动态资源
	app.PartyFunc("/home", func(home iris.Party) {
		home.Get("/multidata", views.HomeMultiData)
	})

	app.Run(iris.Addr(":8081"))
}
