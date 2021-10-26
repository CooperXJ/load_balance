package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"load_balance/loadbalance/controller/common"
	"load_balance/loadbalance/service"
)

func GetRandomServer(c *gin.Context) {
	val := c.Query("serviceName")
	server, err := service.Random(val)
	if err != nil {
		common.ResponseWithMsg(c, common.CodeInternalError, err.Error())
	}

	fmt.Printf("%v服务 ---------------> %v\n", server.ServiceName, server.ServiceEndpoint)
	common.ResponseSuccess(c, server.ServiceEndpoint)
}
