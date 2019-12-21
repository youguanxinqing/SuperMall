package model

// Img ...
type Img struct {
	Image string `json:"path"`
}

// HyperImg 超链接图片
type HyperImg struct {
	Image string `json:"image"`
	Link  string `json:"link"`
	Title string `json:"title"`
}
