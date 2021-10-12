package main

import (
	"math"
	"time"
)

func GetVmLoad(vm *Vm,interval time.Duration,loadChan chan<- float64)  {
	var minCPU  = math.MaxFloat64
	var minMem  = math.MaxFloat64
	var minDisk  = math.MaxFloat64
	var minConn  = math.MaxFloat64

	var maxCPU float64 = 0
	var maxMem float64 = 0
	var maxDisk float64 = 0
	var maxConn float64 = 0

	for  {
		c1 := vm.stats.CPUUsage
		m1 := vm.stats.MemUsage
		d1 := vm.stats.DiskUsage
		s1 := float64(vm.stats.ConnUsage)

		time.Sleep(interval)

		c2 := vm.stats.CPUUsage
		m2 := vm.stats.MemUsage
		d2 := vm.stats.DiskUsage
		s2 := float64(vm.stats.ConnUsage)

		changeCPU := math.Abs(c1-c2)
		changeMem := math.Abs(m1-m2)
		changeDisk := math.Abs(d1-d2)
		changeConn := math.Abs(s1-s2)

		minCPU  = math.Min(math.Min(minCPU,c1),c2)
		minMem  = math.Min(math.Min(minMem,m1),m2)
		minDisk  = math.Min(math.Min(minDisk,d1),d2)
		minConn  = math.Min(math.Min(minConn,s1),s2)
		maxCPU  = math.Max(math.Max(maxCPU,c1),c2)
		maxMem  = math.Max(math.Max(maxMem,m1),m2)
		maxDisk  = math.Max(math.Max(maxDisk,d1),d2)
		maxConn  = math.Max(math.Max(maxConn,s1),s2)

		//正则化
		changeCPU = (changeCPU-minCPU)/(maxCPU-minCPU)
		changeMem = (changeMem-minMem)/(maxMem-minMem)
		changeDisk = (changeDisk-minDisk)/(maxDisk-minDisk)
		changeConn = (changeConn-minConn)/(maxConn-minConn)

		//fmt.Printf("cpu=%v mem=%v disk=%v conn=%v\n",changeCPU,changeMem,changeDisk,changeConn)

		w := changeConn+changeMem+changeDisk+changeConn
		WCPU := changeConn/w
		WMem := changeMem/w
		WDisk := changeDisk/w
		WConn := changeConn/w

		load:=WCPU*c2+WMem*m2+WDisk*d2+WConn*s2
		loadChan<-load
	}
}
