package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main(){

	// 1。原始的http服务
	//s := web.NewService(web.Address(":8081")) // imporant  冒号。  ":8081"
	//s.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("ok"))
	//})
	//s.Run()

	// 2。引入 gin 框架实现http服务
	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/hello", func(context *gin.Context) {
		context.String(200, "hello moto")
	})
	server := web.NewService(web.Address(":8081"), web.Handler(ginRouter))
	server.Run()


}