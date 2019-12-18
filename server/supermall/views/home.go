package views

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
			"banner":    banner,
			"dKeyword":  dKeyword,
			"keywords":  keywords,
			"recommend": recommend,
		},
	})
}
