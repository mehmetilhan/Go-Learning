package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go test(ch)

	time.Sleep(time.Second * 1)
	fmt.Printf("Send Value: %t\n", true)
	ch <- true

	time.Sleep((time.Second * 3))
}

func test(ch chan bool) {
	var value = <-ch
	fmt.Printf("Received Value: %t\n", value)
	fmt.Println("Doing some work")
}

//Unbufferedbir kanal, gönderme ve alma gerçekleştiği anda iki goroutin arasında
//bir değiş tokuşun gerçekleştirildiğini garanti eder. Buffered bir kanalın böyle bir garantisi yoktur.
