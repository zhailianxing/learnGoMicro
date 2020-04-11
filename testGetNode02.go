package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	plguinsHttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)



/* 功能：向consul(地址:127.0.0.1:8500)中查找 指定 服务的信息。
即：1. 查找一个名字为"prodService"的服务，返回这个服务的所有节点(ip,port)  2.使用selector模块返回一个随机的服务节点
*/
func main() {
	// 1. 新建一个cousul客户端
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// 2.通过 client/selector包 获取指定服务的信息。返回值是：[]*Registry.Service
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin), // 设置获取服务节点策略
		)
	// 同样可以获取指定服务的信息
	//next, err :=  mySelector.Select("product2Servcie")
	//node, err := next()
	//fmt.Println(node)

	CallApi2(mySelector)

}

// 使用 go-plugins/client/http 包的封装函数，进行服务调用
func CallApi2( s selector.Selector) {
	// plguinsHttp
	plguinsHttpClient := plguinsHttp.NewClient(
		client.Selector(s), // 这里的client是go-micro下的。和go-plugins下的client是不一样的
		client.ContentType("application/json"),
		)

	// 第一个参数： 请求的服务名字
	// 第二个参数endPoint: http协议用的是 path； grpc用的是 函数名；
	// 第三个参数： 请求参数
	req := plguinsHttpClient.NewRequest("product2Servcie", "/v1/getProduct", nil)
	var resp map[string]interface{} // 根据接口返回的结果 定义类型
	err := plguinsHttpClient.Call(context.Background(), req, &resp) //resp要传指针进去
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resp:", resp)

}