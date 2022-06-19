package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	//ch <- 1 bu satır acık olursa kanal kapasitesinden cok kanala goroutine gonderdıgımızden uygulama deadlock hatası verir
	fmt.Println("Sending value to channel complete")
	val := <-ch
	//val = <-ch bu satır acık olursa kanal kapasitesinden cok kanala goroutine gonderdıgımızden uygulama deadlock hatası verir
	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)
}
