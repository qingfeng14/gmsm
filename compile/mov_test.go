package main

import (
	"fmt"
	"math"
	"sync"
	"testing"
)

func BenchmarkMovR(b *testing.B) {
	x := uint64(1)<<59 + 589345834657
	for i := 0; i < b.N; i++ {
		_ = uint32(x >> 57)
	}
}

func BenchmarkMovRR(b *testing.B) {
	x := uint64(1)<<59 + 589345834657
	for i := 0; i < b.N; i++ {
		_ = uint32(x>>32) >> 25
	}
}

func TestMove(t *testing.T) {
	ch := make(chan int, 266450000)
	var wg sync.WaitGroup
	for i := uint64(0); i <= math.MaxUint64; i++ {
		wg.Add(1)
		ch <- 1
		go func(i uint64) {
			defer func() {
				wg.Done()
				<-ch
			}()
			if i%266450000 == 0 {
				fmt.Println("Now", i)
			}

			j := (uint32(i >> 32)) >> 25
			k := uint32(i >> 57)
			if j != k {
				fmt.Println(i, j, k)
				t.Error(i)
			}
		}(i)
	}
}
