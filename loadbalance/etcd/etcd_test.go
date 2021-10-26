package etcd

import (
	"fmt"
	"load_balance/loadbalance/model"
	"testing"
)

func TestPutService(t *testing.T) {
	config := &model.EtcdConfig{
		EtcdEndpoints:      []string{"https://172.20.4.90:22379", "https://172.20.4.91:22379", "https://172.20.4.92:22379"},
		TlsDisabled:        false,
		EtcdTLSCertPath:    "..\\tls\\cert.pem",
		EtcdTLSCertKeyPath: "..\\tls\\key.pem",
		EtcdTLSCAPath:      "..\\tls\\ca.pem",
	}

	err := Init(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	//defMap:=&model.ServiceDefMap{
	//	ServiceEndpoint: "172.23.27.115:1009",
	//	ServiceName:     "cpu",
	//	ServiceHost:     "172.23.27.115",
	//}
	//
	//err = PutService(defMap)
	//if err != nil {
	//	fmt.Println("无法注册服务")
	//}

	service, err := GetService("/cpu")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(service)

	// del
	//err = DelService(client, "cpu")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}
