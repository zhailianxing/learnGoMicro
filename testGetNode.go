package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
)



/* 功能：向consul(地址:127.0.0.1:8500)中查找 指定 服务的信息。
即：1. 查找一个名字为"prodService"的服务，返回这个服务的所有节点(ip,port)  2.使用selector模块返回一个随机的服务节点
*/
func main() {
	// 1. 新建一个cousul客户端
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// 2.获取指定服务的信息。返回值是：[]*Registry.Service
	prodService, err := consulReg.GetService("product2Servcie")
	if err != nil {
		log.Fatal(err) //内部调用os.Exit(1)
	}

	//3.使用client模块下的selector包，随机获取一个节点
	next :=  selector.Random(prodService)
	//next是一个func.  type Next func() (*registry.Node, error)
	node, err := next()
	if err != nil {
		log.Fatal(err) //内部调用os.Exit(1)
	}
	fmt.Println(node)

	//3.调用此服务(product2Servcie)的接口;使用的最原始的http访问
	d,_ := CallApiGet(node.Address, "/v1/getProduct", "")
	fmt.Println("data:",d)

}

// 只是为了演示功能，这种写法不可取。
func CallApiGet( addrs, path, body string) (string, error) {
	req, _ := http.NewRequest("GET", "http://"+addrs+path, nil)
	fmt.Println("path:", addrs+path)
	client := http.Client{}
	resp, err := client.Do(req)
	defer  resp.Body.Close()
	if err != nil {
		return "", err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil

}