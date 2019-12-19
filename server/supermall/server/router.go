package server

import (
	"github.com/kataras/iris/v12"

	"supermall/api"
	"supermall/middleware"
)

// NewApplication ...
func NewApplication() *iris.Application {
	app := iris.Default()

	// 中间件设置
	app.Use(middleware.AccessControlAllowOrigin)

	/* 注册路由 */

	// 静态资源
	app.Get("/static/{filepath:path}", api.StaticResource)

	// 动态资源
	app.PartyFunc("/home", func(home iris.Party) {
		home.Get("/multidata", api.HomeMultiData)
	})

	return app
}
