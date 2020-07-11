package main

import (
	"carrotCloud/rout"
	"carrotCloud/config"
	"fmt"
)

func main() {
	// 路由设置
	router := rout.Router()

	// 启动服务并且监听接口
	err := router.Run(config.UploadServiceHost)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s\n", err.Error())
	}


}

