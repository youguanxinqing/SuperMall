package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
	fileNoExt := trimExt(filename)
	// 从配置文件中查找图片对应的 url, 如果不存在, 默认为首页
	fmt.Println(fileNoExt)
	infos, ok := config.ImgURLMap[fileNoExt]
	var title, urlPath string
	if !ok {
		urlPath = "/"
		title = ""
	} else {
		urlPath = infos[config.URLPATH]
		title = infos[config.TITLE]
	}

	return model.HyperImg{
		Image: NormalPathToURL(path),
		Link:  CompleteURLPath(urlPath),
		Title: title,
	}
}

// 移除文件的扩展名
// eg:
// zhong.jpg => zhong
func trimExt(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	return strings.Trim(filename, ext)
}
