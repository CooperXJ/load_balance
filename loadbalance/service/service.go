package service

import (
	"load_balance/loadbalance/etcd"
	"load_balance/loadbalance/model"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(serviceName string) (*model.ServiceDefMap, error) {
	services, err := etcd.GetService(serviceName)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, nil
	}

	return &services[rand.Intn(len(services))], nil
}
