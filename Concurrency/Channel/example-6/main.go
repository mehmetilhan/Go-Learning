package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 3
	ch <- 4

	close(ch)

	go sum(ch)
	time.Sleep(time.Second * 1)
}

func sum(ch chan int) {
	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Printf("Sum : %d\n", sum)
}
