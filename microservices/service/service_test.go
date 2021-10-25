package service

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func Test_randStringRunes(t *testing.T) {
	fmt.Println(randStringRunes(8))
}

func TestRand(t *testing.T) {
	cur := time.Now()
	temp := rand.Float64()
	ticker := time.NewTicker(time.Second * 10)

loop:
	for {
		select {
		case <-ticker.C:
			break loop
		default:
			temp = math.Sqrt(math.Sqrt(temp))
		}
	}

	fmt.Println(rand.Float64())
	fmt.Println(time.Now().Sub(cur))
}
