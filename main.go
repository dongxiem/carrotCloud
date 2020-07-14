package main

import (
	"carrotCloud/config"
	"carrotCloud/route"
	"fmt"
)

func main() {
	// 路由设置
	router := route.Router()

	// 启动服务并且监听接口
	err := router.Run(config.UploadServiceHost)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s\n", err.Error())
	}

}
