package service

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MemType() {
	var a []string
	//ticker := time.NewTicker(time.Second * 2)
	str := randStringRunes(8)
	for i := 1; i < 1000; i++ {
		a = append(a, str)
	}

	//<-ticker.C
	fmt.Println(a[len(a)-1])
}

func CPUType() {
	temp := rand.Float64()
	//	ticker := time.NewTicker(time.Millisecond * 200)
	//
	//loop:
	//	for {
	//		select {
	//		case <-ticker.C:
	//			break loop
	//		default:
	//			temp = math.Sqrt(math.Sqrt(temp))
	//		}
	//	}

	for i := 0; i < 10; i++ {
		temp = math.Sqrt(math.Sqrt(temp))
	}
}
