package main

import (
	"fmt"
	"math"
)

func Pipeline(in <-chan uint8, out chan<- float64) {
	for val := range in {
		out <- math.Pow(float64(val), 3)
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go Pipeline(in, out)

	go func() {
		for _, val := range []uint8{1, 2, 3, 4, 5} {
			in <- val
		}
		close(in)
	}()

	for result := range out {
		fmt.Println(result)
	}
}
