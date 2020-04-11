package main


//大坑啊：github.com/micro/go-micro/registry/consul用作注册中心，在1.14.0版本之后删除了，需要用之前的micro版本. 或者换成etcd作为注册中心
//"github.com/micro/go-micro/registry/consul"

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-micro/web"
)

func main(){
	//1.使用 go-plugins 中内置的consul
	//2. 127.0.0.1:8500是consul的地址。
	consuReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/hello", func(context *gin.Context) {
		context.String(200, "hello moto")
	})

	//向consuReg所代表的cousul里，注册名字为"prodServcie"的服务，此服务的地址是"127.0.0.1:8081"
	server := web.NewService(
		web.Registry(consuReg),
		web.Name("indexServcie"),
		//web.Address(":8081"), // 即：127.0.0.1:8081的简写
		web.Handler(ginRouter),
		)

	server.Init() // 内部基础cmd命令解析命令行参数。 参数是：server_name，server_address
	server.Run()

}
//命令行执行：
// 注册第一个服务：go run main.go --server_address :8081
// 注册第二个服务：go run main.go --server_address :8082

//最后：所有的服务信息应该放在配置文件中或者配置服务中(consul/etcd等)，而不是从命令行直接写，这样不易维护。