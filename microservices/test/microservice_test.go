package test

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func TestMem(t *testing.T) {
	var a []string

	for i := 1; i < 31000000; i++ {
		a = append(a, "test")
	}
	log.Println("end---")
	PrintMemUsage()
	for {
		runtime.GC() //调用 runtime.GC() 进行手动触发GC进行内存回收
		PrintMemUsage()
		log.Println(a[len(a)-1])
		time.Sleep(time.Second * 2)
	}
	time.Sleep(time.Second * 100)
	log.Println("test memory end...", len(a))
}
