package db

import (
	"io/ioutil"
	"os"
)

// Img 图片
type Img struct {
	Path string `json:"path"`
}

// bindImgPath 绑定图片路径
func (img *Img) bindImgPath(path string) {
	img.Path = path
}

// BindImgsByDir 传入目录路径，绑定目录下的图片文件
func BindImgsByDir(dir string) ([]Img, error) {
	var imgs []Img
	// 打开目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	sep := string(os.PathSeparator)
	for _, fObj := range files {
		// 略过目录
		if fObj.IsDir() {
			continue
		}
		imgs = append(imgs, Img{dir + sep + fObj.Name()})
	}
	return imgs, nil
}
