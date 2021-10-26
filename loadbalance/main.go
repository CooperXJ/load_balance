package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"load_balance/loadbalance/controller"
	"load_balance/loadbalance/etcd"
	"load_balance/loadbalance/model"
	"net/http"
)

func main() {
	config := &model.EtcdConfig{
		EtcdEndpoints:      []string{"https://172.20.4.90:22379", "https://172.20.4.91:22379", "https://172.20.4.92:22379"},
		TlsDisabled:        false,
		EtcdTLSCertPath:    "./tls/cert.pem",
		EtcdTLSCertKeyPath: "./tls/key.pem",
		EtcdTLSCAPath:      "./tls/ca.pem",
	}

	err := etcd.Init(config)
	if err != nil {
		fmt.Printf("初始化etcd失败\n %v\n", err.Error())
	}
	defer etcd.Close()

	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	monGroup := r.Group("/select")
	monGroup.GET("/random", controller.GetRandomServer)

	r.Run(":8090")
}
