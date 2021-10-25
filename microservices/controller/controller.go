package controller

import (
	"github.com/gin-gonic/gin"
	"load_balance/microservices/controller/common"
	"load_balance/microservices/service"
)

func MemTypeHandler(c *gin.Context) {
	service.MemType()
	common.ResponseSuccess(c, common.CodeSuccess)
}

func CPUTypeHandler(c *gin.Context) {
	service.CPUType()
	common.ResponseSuccess(c, common.CodeSuccess)
}
