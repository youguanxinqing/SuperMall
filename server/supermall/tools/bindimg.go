package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"supermall/config"
	"supermall/model"
)

// BindHyperImgsByDir 传入目录路径，绑定目录下的图片文件
func BindHyperImgsByDir(dir string) ([]model.HyperImg, error) {
	var imgs []model.HyperImg
	// 打开目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fObj := range files {
		// 略过目录
		if fObj.IsDir() {
			continue
		}
		imgs = append(imgs, generateHyperImg(dir, fObj.Name()))
	}
	return imgs, nil
}

// generateHyperImg 生成 HyperImg
func generateHyperImg(dir, filename string) model.HyperImg {
	path := dir + string(os.PathSeparator) + filename

	fileNoExt := filepath.Base(filename)
	// 从配置文件中查找图片对应的 url, 如果不存在, 默认为首页
	url, ok := config.ImgURLMap[fileNoExt]
	var urlPath string
	if !ok {
		urlPath = "/"
	} else {
		urlPath = url
	}

	link := fmt.Sprintf(
		"http://%s:%d%s",
		config.ServerHost, config.ServerPort, urlPath,
	)

	return model.HyperImg{
		Image: path,
		Link:  link,
	}
}
