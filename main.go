package main

import (
	"todo-list-server/conf"
	"todo-list-server/routes"
)

func main() {
	// 读取配置
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run()
}
