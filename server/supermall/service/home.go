package service

import (
	"supermall/config"
	"supermall/tools"

	"github.com/kataras/iris/v12"
)

// generateImgInfoByDir 组装轮播图相关数据
func generateImgInfoByDir(dir string) (res []iris.Map) {
	imgs, err := tools.BindHyperImgsByDir(dir)
	if err != nil {
		return []iris.Map{}
	}

	for _, img := range imgs {
		res = append(res, tools.StructToMap(img))
	}
	return
}

// HomeMultiDataService 组合首页数据
func HomeMultiDataService() ([]iris.Map, []iris.Map, []iris.Map, []iris.Map) {
	swiperImgDir := tools.ConfigToNoramlPath(config.SwiperImgDIr)
	recommandImgDir := tools.ConfigToNoramlPath(config.RecommandImgDir)
	return generateImgInfoByDir(swiperImgDir), // 轮播图
		[]iris.Map{},
		[]iris.Map{},
		generateImgInfoByDir(recommandImgDir) // 推荐
}
