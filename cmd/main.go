package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-basic/src/views"
	"net/http"
)

func main() {
	mainrouter := chi.NewRouter()
	// 创建二级路由，用户路由
	userouter := chi.NewRouter()
	userouter.Post("/userinfo", views.Useradd)

	// 将二级路由挂载到一级路由中
	mainrouter.Mount("/user", userouter)
	// 启动web服务
	err := http.ListenAndServe(":3000", mainrouter)
	if err != nil {
		panic("服务启动异常")
	} else {
		fmt.Println("服务启动正常,端口号为:3000")
	}
}
