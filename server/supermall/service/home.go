package service

import (
	"supermall/config"
	"supermall/tools"

	"github.com/kataras/iris/v12"
)

// GetSwiperImg 组装轮播图相关数据
func makeSwiperImg(dir string) (res []iris.Map) {
	imgs, err := tools.BindHyperImgsByDir(dir)
	if err != nil {
		return []iris.Map{}
	}

	for _, img := range imgs {
		res = append(res, iris.Map{
			"image": tools.NormalPathToURL(img.Image),
			"link":  img.Link,
		})
	}
	return
}

// HomeMultiDataService 组合首页数据
func HomeMultiDataService() ([]iris.Map, []iris.Map, []iris.Map, []iris.Map) {
	swiperImgDir := tools.ConfigToNoramlPath(config.SwiperImgDIr)
	return makeSwiperImg(swiperImgDir), []iris.Map{}, []iris.Map{}, []iris.Map{}
}
