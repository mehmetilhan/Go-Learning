package main

import "fmt"

//Benzer sınıfların oluşturulmasında kolaylık sağlar.

func main() {

	bmw, _ := getCar("bmw")
	mercedes, _ := getCar("mercedes")
	fmt.Printf("%s is %d hp", bmw.getName(), bmw.getHP())
	fmt.Println()
	fmt.Printf("%s is %d hp", mercedes.getName(), mercedes.getHP())

}
