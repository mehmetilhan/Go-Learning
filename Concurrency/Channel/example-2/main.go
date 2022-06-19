package main

import (
	"fmt"
	"time"
)

//Bu, arabelleğe alınmamış bir kanalda göndermenin,
//başka bir goroutinde o kanalda bir alma gerçekleşene kadar blok olduğunu gösterir.

func main() {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 5)
}

func send(ch chan int) { //send işlemine giren goroutine bloklanır ve işlem yapmaz
	ch <- 1
	fmt.Println("Sending value to channel complete")
}

func receive(ch chan int) { // ta ki receive işlemi gerçekleştikten sonra goroutine blok kalkar
	time.Sleep(time.Second * 3)
	fmt.Println("Timeout finished")
	_ = <-ch
	return
}
