package main

import (
	"carrotCloud/handler"
)

func main() {
	api := handler.InitRouter()
	api.Run(":5000")
}
