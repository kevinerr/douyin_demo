package main

import (
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/gin-gonic/gin"
)

// @title DouYin API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name hkh
// @BasePath /douyin
func main() { // http://localhost:8080/swagger/index.html

	//从配置文件读入配置
	conf.Init()

	//转载路由 swag init -g common.go
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
