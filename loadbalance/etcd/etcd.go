package etcd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"io/ioutil"
	"load_balance/loadbalance/model"
	"strings"
	"time"
)

var client *clientv3.Client

// Init 获取etcd客户端
func Init(conf *model.EtcdConfig) (err error) {
	var etcdCfg clientv3.Config
	switch conf.TlsDisabled {
	case false:
		cert, err := tls.LoadX509KeyPair(conf.EtcdTLSCertPath, conf.EtcdTLSCertKeyPath)
		if err != nil {
			return err
		}
		caData, err := ioutil.ReadFile(conf.EtcdTLSCAPath)
		if err != nil {
			return err
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caData)
		etcdCfg = clientv3.Config{
			Endpoints: conf.EtcdEndpoints,
			TLS: &tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs:      pool,
			},
			DialTimeout:          time.Second * 5,
			DialKeepAliveTime:    time.Second * 10,
			DialKeepAliveTimeout: time.Second * 10,
			AutoSyncInterval:     time.Minute * 10,
			DialOptions:          []grpc.DialOption{grpc.WithBlock()},
		}
	case true:
		etcdCfg = clientv3.Config{
			Endpoints: conf.EtcdEndpoints,
		}
	}
	client, err = clientv3.New(etcdCfg)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	client.Close()
}

// PutService 注册微服务
func PutService(serviceDefMap *model.ServiceDefMap) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defMap, err := json.Marshal(serviceDefMap)
	if err != nil {
		return err
	}

	// key = serviceName/ip/port
	arr := strings.Split(serviceDefMap.ServiceEndpoint, ":")
	serviceName := fmt.Sprintf("/%s/%s/%s", serviceDefMap.ServiceName, arr[0], arr[1])
	_, err = client.Put(ctx, serviceName, string(defMap))
	if err != nil {
		return err
	}

	cancel()

	return nil
}

func DelService(serviceName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := client.Delete(ctx, serviceName)
	if err != nil {
		return err
	}

	cancel()
	return nil
}

func GetService(servicePrefix string) ([]model.ServiceDefMap, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, servicePrefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	services := make([]model.ServiceDefMap, len(resp.Kvs))
	for i := 0; i < len(services); i++ {
		var service model.ServiceDefMap
		if err = json.Unmarshal(resp.Kvs[i].Value, &service); err != nil {
			return nil, err
		}
		services[i] = service
	}

	cancel()
	return services, nil
}
