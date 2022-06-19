package main

import "fmt"

/*
	chan  :bidirectional channel (Both read and write)
	chan <-  :only writing to channel
	<- chan  :only reading from channel (input channel)
*/

/*func main() { //Only Send Channel
	ch := make(chan int, 3)
	process(ch)
	fmt.Println(<-ch)
}
func process(ch chan<- int) {
	ch <- 2
}*/

func main() { //Only Receive Channel
	ch := make(chan int, 3)
	ch <- 2
	process(ch)
	fmt.Println()
}
func process(ch <-chan int) {
	s := <-ch
	fmt.Println(s)
}
