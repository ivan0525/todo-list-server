package main

import (
	"todo-list-server/conf"
	"todo-list-server/routes"
)

func main() {
	// 读取配置
	conf.Init()
	r := routes.NewRouter()
	// 服务启动端口，默认为8080
	_ = r.Run(conf.HttpPort)
}
