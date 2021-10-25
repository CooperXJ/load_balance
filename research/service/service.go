package service

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"hostMonitor/model"
	"time"
)

// GetCPU 获取CPU信息
func GetCPU() (*model.CPU, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	res := &model.CPU{}
	for _, info := range percent {
		res.Used = info
	}

	return res, nil
}

//GetMem 获取内存信息
func GetMem() (*model.Mem, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	res := &model.Mem{}
	res.Total = memory.Total
	res.Available = memory.Available
	res.Used = memory.Used
	res.UsedPercent = memory.UsedPercent
	return res, nil
}

// GetSumReport 获取总的结果
func GetSumReport() (*model.SumReport, error) {
	CPU, err := GetCPU()
	if err != nil {
		return nil, err
	}
	Mem, err := GetMem()
	if err != nil {
		return nil, err
	}

	sumReport := &model.SumReport{
		CPU:    CPU,
		Memory: Mem,
		Disk:   &model.Disk{},
		NetIO:  &model.NetIO{},
	}

	return sumReport, nil
}
