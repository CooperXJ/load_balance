package controller

import (
	"github.com/gin-gonic/gin"
	"hostMonitor/controller/common"
	"hostMonitor/service"
)

func GetCPUHandler(c *gin.Context) {
	CPU, err := service.GetCPU()
	if err != nil {
		common.ResponseWithMsg(c, common.CodeCPUGetFailed, err.Error())
	}

	common.ResponseSuccess(c, CPU)
}

func GetMemHandler(c *gin.Context) {
	mem, err := service.GetMem()
	if err != nil {
		common.ResponseWithMsg(c, common.CodeMemGetFailed, err.Error())
	}

	common.ResponseSuccess(c, mem)
}

func GetSumReportHandler(c *gin.Context) {
	sumReport, err := service.GetSumReport()
	if err != nil {
		common.ResponseWithMsg(c, common.CodeSumFailed, err.Error())
	}

	common.ResponseSuccess(c, sumReport)
}
