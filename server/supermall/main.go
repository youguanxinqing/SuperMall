package main

import (
	"fmt"
	"supermall/config"
	"supermall/server"

	"github.com/kataras/iris/v12"
)

func main() {
	app := server.NewApplication()
	addr := fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	app.Run(iris.Addr(addr))
}
