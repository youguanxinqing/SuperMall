package config

/*
 * 路径均采用 ‘|’(竖线) 分割, 以动态的方式进行路径拼接，实现系统的跨平台操作
 *
**/

// SwiperImgDIr 轮播图目录
const SwiperImgDIr = "static|img|swiper"

// RecommandImgDir 推荐图目录
const RecommandImgDir = "static|img|recommand"

// ImgURLMap 图片关系映射
var ImgURLMap = map[string][]string{
	// title, link
	"swiper01":    []string{"", "/"},
	"swiper02":    []string{"", "/"},
	"swiper03":    []string{"", "/"},
	"swiper04":    []string{"", "/"},
	"recommand01": []string{"推荐1", "/"},
	"recommand02": []string{"推荐2", "/"},
	"recommand03": []string{"推荐3", "/"},
	"recommand04": []string{"推荐4", "/"},
}

const (
	// TITLE 标题
	TITLE = 0
	// URLPATH url path 值
	URLPATH = 1
)
