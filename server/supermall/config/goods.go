package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// GoodsData 商品数据
var GoodsData map[string]interface{}

// 家在 Goods.json 文件
func init() {
	f, err := os.Open(goodsJSONPath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	dataBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	json.Unmarshal(dataBytes, &GoodsData)
}
