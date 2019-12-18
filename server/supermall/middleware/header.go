package middleware

import "github.com/kataras/iris/v12"

// AccessControlAllowOrigin 允许跨域访问
func AccessControlAllowOrigin(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}
