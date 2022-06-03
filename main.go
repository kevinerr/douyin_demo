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

//func main() {
//	fmt.Println("OSS Go SDK Version: ", oss.Version)
//	endpoint := "oss-cn-shanghai.aliyuncs.com"
//	accessKey := "LTAI4GEi2cat7zLt37PSrixz"
//	secretKey := "6iNRN9bdVJKC5gyRJruWnIHlWrApH2"
//	client, err := oss.New(endpoint, accessKey, secretKey, oss.Timeout(10, 120))
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//	// 选择桶
//	bucket, err := client.Bucket("paper-boot")
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//	// 上传文件。
//	err = bucket.PutObjectFromFile("test/1.png", "/Users/xietingyu/Desktop/1.png")
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//}
