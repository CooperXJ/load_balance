package main

import (
	"fmt"
	"testing"
	"time"
)

// 测试RandFloat
func TestRandFloats(t *testing.T) {
	fmt.Println(randFloats(0.2,1.5))
}

func TestRandInts(t *testing.T) {
	fmt.Println(randInts(1,10))
}

func TestStart(t *testing.T) {
	syncChan:=make(chan struct{},1)

	vm:=&Vm{
		Job: "oder",
		CPU:       4,
		Mem:       8*1024,
		Disk:      500,
		Conn:      1000,
	}

	stats:= VmConfig{
		CPUUsageMax:  3,
		CPUUsageMin:  1,
		MemUsageMax:  7,
		MemUsageMin:  1,
		DiskUsageMax: 450,
		DiskUsageMin: 5,
		ConnCountMax: 950,
		ConnCountMin: 500,
	}
	stopChan:=make(chan struct{},1)
	vm.Start(stats,stopChan,500*time.Millisecond)

	go func(vm *Vm) {
		for {
			time.Sleep(time.Second)
			vm.GenerateVmReport()
			fmt.Println(&vm.report)
		}
	}(vm)

	<-syncChan
}

func TestGetVmLoad(t *testing.T) {
	syncChan:=make(chan struct{},1)
	loadChan:=make(chan float64,10)
	vm:=&Vm{
		CPU:       4,
		Mem:       8*1024,
		Disk:      500,
		Conn:      1000,
	}

	stats:= VmConfig{
		CPUUsageMax:  3,
		CPUUsageMin:  1,
		MemUsageMax:  7,
		MemUsageMin:  1,
		DiskUsageMax: 450,
		DiskUsageMin: 5,
		ConnCountMax: 950,
		ConnCountMin: 500,
	}
	stopChan:=make(chan struct{},1)
	vm.Start(stats,stopChan,500*time.Millisecond)

	go GetVmLoad(vm,2*time.Second,loadChan)

	for load:=range loadChan{
		fmt.Println(load)
	}

	<-syncChan

}