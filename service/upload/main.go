package main

import (
	cfg "carrotCloud/config"
	"carrotCloud/route"
	"fmt"
)

func main() {
	fmt.Printf("上传服务启动中，开始监听监听[%s]...\n", cfg.UploadServiceHost)
	// 路由设置
	router := route.Router()
	// 启动服务并且监听接口
	err := router.Run(cfg.UploadServiceHost)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s\n", err.Error())
	}
}
