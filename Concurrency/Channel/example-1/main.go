package main

import (
	"fmt"
	"time"
)

//varsayılan olarak unbuffered bir channel oluşuturldugu için
//ilk send işleminden sonra send block lanır
//recieve calısana kadar bloklu kalır
func main() {
	ch := make(chan int)

	fmt.Println("Sending value to channel")
	go send(ch)

	fmt.Println("Receiving from channel")
	go receive(ch)

	time.Sleep(time.Second * 1) // bu sleep işlemi goroutine cevabı gelene kadar uygulama kapanmasın diye var
}

func send(ch chan int) {
	ch <- 1
}

func receive(ch chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)
}
