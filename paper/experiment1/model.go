package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Vm 虚拟机   主要包含CPU使用率、内存使用率、磁盘使用率、当前服务器连接数
type (
	Vm struct {
		Name string
		Job  Job
		CPU  float64
		Mem  float64
		Disk float64
		Conn int

		config VmConfig
		stats  VmStats
		report VmReport
	}

	// VmConfig 虚拟机运行的配置（设置各项的max、min）
	VmConfig struct {
		CPUUsageMax  float64
		CPUUsageMin  float64
		MemUsageMax  float64
		MemUsageMin  float64
		DiskUsageMax float64
		DiskUsageMin float64
		ConnCountMax int
		ConnCountMin int
	}

	// VmStats 虚拟机运行状态
	VmStats struct {
		CPUUsage  float64
		MemUsage  float64
		DiskUsage float64
		ConnUsage int
	}

	// VmReport 虚拟机运行报表
	VmReport struct {
		CPUUtility  float64
		MemUtility  float64
		DiskUtility float64
		ConnUtility float64
	}
)

func (vm *Vm) ControlCPUUsage(min, max float64) error {
	if min < 0 && max > vm.CPU {
		return errors.New("CPU使用量设置有问题")
	}
	vm.stats.CPUUsage = randFloats(min, max)
	return nil
}

func (vm *Vm) ControlMemUsage(min, max float64) error {
	if min < 0 && max > vm.Mem {
		return errors.New("内存使用量设置有问题")
	}
	vm.stats.MemUsage = randFloats(min, max)
	return nil
}

func (vm *Vm) ControlDiskUsage(min, max float64) error {
	if min < 0 && max > vm.Disk {
		return errors.New("磁盘使用量设置有问题")
	}
	vm.stats.DiskUsage = randFloats(min, max)
	return nil
}

func (vm *Vm) ControlConnCount(min, max int) error {
	if min < 0 && max > vm.Conn {
		return errors.New("连接数设置有问题")
	}
	vm.stats.ConnUsage = randInts(min, max)
	return nil
}

func (vm *Vm) Start(config VmConfig, ctx context.Context, interval time.Duration) {
	go func(status VmConfig, ctx context.Context) {
		ticker := time.NewTicker(interval)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				ticker.Stop()
			default:
				err := vm.ControlCPUUsage(status.CPUUsageMin, status.CPUUsageMax)
				if err != nil {
					fmt.Println(err.Error())
				}
				err = vm.ControlMemUsage(status.MemUsageMin, status.MemUsageMax)
				if err != nil {
					fmt.Println(err.Error())
				}
				err = vm.ControlDiskUsage(status.DiskUsageMin, status.DiskUsageMax)
				if err != nil {
					fmt.Println(err.Error())
				}
				err = vm.ControlConnCount(status.ConnCountMin, status.ConnCountMax)
				if err != nil {
					fmt.Println(err.Error())
				}

				vm.GenerateVmReport()
				busyPolicy(&vm.report, vm.Name)
			}
		}
	}(config, ctx)
}

func (vm *Vm) GenerateVmReport() {
	vm.report.CPUUtility = vm.stats.CPUUsage / vm.CPU
	vm.report.MemUtility = vm.stats.MemUsage / vm.Mem
	vm.report.DiskUtility = vm.stats.DiskUsage / vm.Disk
	vm.report.ConnUtility = float64(vm.stats.ConnUsage) / float64(vm.Conn)
}

func (vm *Vm) String() string {
	return fmt.Sprintf(
		"vm job:\n job = %v"+
			"vm config:\n cpu = %v,memory = %v,disk = %v,connection = %v\n"+
			"vm stats:\n cpu = %v,memory = %v,disk = %v,connection = %v\n"+
			"vm report:\n cpu = %v,memory = %v,disk = %v,connection = %v\n",
		vm.Job,
		vm.CPU, vm.Mem, vm.Disk, vm.Conn,
		vm.stats.CPUUsage, vm.stats.MemUsage, vm.stats.DiskUsage, vm.stats.ConnUsage,
		vm.report.CPUUtility, vm.report.MemUtility, vm.report.DiskUtility, vm.report.ConnUtility)
}

func (stats *VmStats) String() string {
	return fmt.Sprintf("vm stats: cpu=%v,memory= %v,disk= %v, connection= %v\n", stats.CPUUsage, stats.MemUsage, stats.DiskUsage, stats.ConnUsage)
}

func (report *VmReport) String() string {
	return fmt.Sprintf("vm report: cpu=%v, memory= %v, disk = %v, connection= %v\n", report.CPUUtility, report.MemUtility, report.DiskUtility, report.ConnUtility)
}

func busyPolicy(report *VmReport, name string) {
	if report.CPUUtility > 0.9 {
		fmt.Printf("%v主机,CPU过高:%v,将睡眠5s\n", name, report.CPUUtility)
		time.Sleep(time.Second * 5)
		return
	} else if report.CPUUtility > 0.85 {
		fmt.Printf("%v主机CPU过高:%v,将睡眠1s\n", name, report.CPUUtility)
		time.Sleep(time.Second)
		return
	}

	if report.MemUtility > 0.95 {
		fmt.Printf("%v主机内存过高:%v,将睡眠3s\n", name, report.MemUtility)
		time.Sleep(time.Second * 3)
		return
	}

	if report.DiskUtility > 0.9 {
		fmt.Printf("%v主机磁盘占用过高:%v,将睡眠5s\n", name, report.DiskUtility)
		time.Sleep(time.Second * 5)
		return
	}

	if report.ConnUtility == 1 {
		fmt.Printf("%v主机连接数过多:%v,将睡眠5s\n", name, report.ConnUtility)
		time.Sleep(time.Second * 5)
		return
	} else if report.ConnUtility > 0.8 {
		fmt.Printf("%v主机连接数过多:%v,将睡眠2s\n", name, report.ConnUtility)
		time.Sleep(time.Second * 2)
		return
	}
}

func randFloats(min, max float64) float64 {
	rand.Seed(time.Now().Unix())
	return min + rand.Float64()*(max-min)
}

func randInts(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
