package api

import (
	"supermall/service"

	"github.com/kataras/iris/v12"
)

// HomeMultiData 返回首页数据
func HomeMultiData(ctx iris.Context) {

	banner, dKeyword, keywords, recommend := service.HomeMultiDataService()

	ctx.JSON(iris.Map{
		"status": "ok",
		"data": iris.Map{
			"banners":    banner,
			"dKeyword":   dKeyword,
			"keywords":   keywords,
			"recommands": recommend,
		},
	})
}
