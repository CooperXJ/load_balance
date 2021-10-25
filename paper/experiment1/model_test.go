package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 测试RandFloat
func TestRandFloats(t *testing.T) {
	fmt.Println(randFloats(0.2, 1.5))
}

func TestRandInts(t *testing.T) {
	fmt.Println(randInts(1, 10))
}

func TestStart(t *testing.T) {
	syncChan := make(chan struct{}, 1)

	vm := &Vm{
		Name: "vm1",
		Job:  "oder",
		CPU:  4,
		Mem:  8 * 1024,
		Disk: 500,
		Conn: 1000,
	}

	stats := VmConfig{
		CPUUsageMax:  3,
		CPUUsageMin:  1,
		MemUsageMax:  7,
		MemUsageMin:  1,
		DiskUsageMax: 450,
		DiskUsageMin: 5,
		ConnCountMax: 950,
		ConnCountMin: 500,
	}
	ctx, cancel := context.WithCancel(context.Background())
	vm.Start(stats, ctx, 500*time.Millisecond)

	go func(vm *Vm) {
		for {
			time.Sleep(time.Second)
			vm.GenerateVmReport()
			fmt.Println(&vm.report)
		}
	}(vm)

	time.Sleep(time.Second * 10)
	cancel()
	<-syncChan
}

func TestGetVmLoad(t *testing.T) {
	syncChan := make(chan struct{}, 1)
	loadChan := make(chan float64, 10)
	vm := &Vm{
		Name: "vm1",
		Job:  "oder",
		CPU:  4,
		Mem:  8 * 1024,
		Disk: 500,
		Conn: 1000,
	}

	stats := VmConfig{
		CPUUsageMax:  3,
		CPUUsageMin:  1,
		MemUsageMax:  7,
		MemUsageMin:  1,
		DiskUsageMax: 450,
		DiskUsageMin: 5,
		ConnCountMax: 950,
		ConnCountMin: 500,
	}

	ctx, cancel := context.WithCancel(context.Background())
	vm.Start(stats, ctx, 500*time.Millisecond)

	go GetVmLoad(vm, 2*time.Second, loadChan)

	for load := range loadChan {
		fmt.Println(load)
	}

	time.Sleep(time.Second * 10)
	cancel()
	<-syncChan

}
