package api

import (
	"io/ioutil"
	"os"

	"github.com/kataras/iris/v12"
)

// StaticResource 静态资源视图
func StaticResource(ctx iris.Context) {
	path := ctx.Params().GetString("filepath")

	prefix := "static" + string(os.PathSeparator)
	file, err := os.Open(prefix + path)
	defer file.Close()

	if err != nil {
		ctx.NotFound()
	} else {
		bs, _ := ioutil.ReadAll(file)
		ctx.Write(bs)
	}
}
