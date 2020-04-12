package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-micro/web"
	"myGoMicro/prodcutService"
)

func main(){
	//1.使用 go-plugins 中内置的consul
	//2. 127.0.0.1:8500是consul的地址。
	consuReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	ginRouter := gin.Default()
	v1Group := ginRouter.Group("v1")
	//使用大括号代码块，仅仅是为了方便看代码。很好
	{
		// v1/getProduct
		v1Group.Handle("GET", "/getProduct", func(context *gin.Context) {
			//context.String(200, "hello moto")
			ret := prodcutService.NewProducts(10)
			context.JSON(200, ret)
		})

		v1Group.Handle("POST", "/getProduct", func(context *gin.Context) {
			//使用gin的方法
			type Req struct {
				Size int `form:"size"` // 客户端使用form表单传递。size字段
			}
			req := Req{}
			err := context.Bind(&req) //使用gin的Bind函数
			if err != nil || req.Size <= 0 {
				req.Size = 3
			}
			ret := prodcutService.NewProducts(req.Size)

			context.JSON(200, gin.H{
				"data":ret,
			})
		})
	}

	//向consuReg所代表的cousul里，注册名字为"prodServcie"的服务，此服务的地址是"127.0.0.1:8081"
	//notice: 一旦此程序退出,就会执行反注册，也就是将 product2Servcie服务配置 从consul中删除
	server := web.NewService(
		web.Registry(consuReg),
		web.Name("product2Servcie"),
		//web.Address(":8082"), // 即：127.0.0.1:8081的简写
		web.Handler(ginRouter),
		)

	server.Init()
	server.Run()

}