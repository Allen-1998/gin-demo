package main

import (
	"gin-demo/routers"
	"gin-demo/routers/user"
)

func main() {
	// 注册路由
	r := routers.Init(user.Router)
	// 监听端口，默认在8080
	r.Run()
}
