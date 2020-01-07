package service

import (
	"fmt"
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

// HomeDataService 组装首页数组
func HomeDataService(typeValue string, page int) (map[string]interface{}, error) {
	data, ok := config.GoodsData[typeValue]
	if !ok {
		return nil, fmt.Errorf("no %s data", typeValue)
	}
	// 对 Goods.json 格式判断
	if dataWithType, ok := data.(map[string]interface{}); ok {
		return map[string]interface{}{
			typeValue: dataWithType, // 流行
		}, nil
	}
	return nil, fmt.Errorf("file Goods.json error formater")
}
