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

// HomeData ...
func HomeData(ctx iris.Context) {
	typeValue := ctx.URLParam("type")
	page, _ := ctx.URLParamInt("page")
	data, err := service.HomeDataService(typeValue, page)
	if err != nil {
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(iris.Map{
		"data": data,
	})
}
