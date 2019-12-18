package service

import (
	"supermall/config"
	"supermall/db"
	"supermall/tools"

	"github.com/kataras/iris/v12"
)

// GetSwiperImg 组装轮播图相关数据
func makeSwiperImg(dir string) (res []iris.Map) {
	imgs, err := db.BindImgsByDir(dir)
	if err != nil {
		return []iris.Map{}
	}

	for _, img := range imgs {
		res = append(res, iris.Map{
			"link": tools.NormalPathToURL(img.Path),
		})
	}
	return
}

// HomeMultiDataService 组合首页数据
func HomeMultiDataService() ([]iris.Map, []iris.Map, []iris.Map, []iris.Map) {
	swiperImgDir := tools.ConfigToNoramlPath(config.SwiperImgDIr)
	return makeSwiperImg(swiperImgDir), []iris.Map{}, []iris.Map{}, []iris.Map{}
}
