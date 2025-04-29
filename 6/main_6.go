package main

import (
	"fmt"
	"math/rand"
)

func Generator(ch chan<- int, count int) {
	for i := 0; i < count; i++ {
		ch <- rand.Intn(10)
	}
	close(ch)
}

func main() {
	const countNums = 0
	ch := make(chan int)
	go Generator(ch, countNums)

	for num := range ch {
		fmt.Println(num)
	}
}
