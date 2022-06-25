package main

import "fmt"

//Defer  work  "LIFO" mechanism
//It works with values when defining functions
func main() {

	i := 0
	defer ShowMe(i) // 2. second work

	i = 5

	defer ShowMe(i) // 1. first work
}

func ShowMe(i int) {
	fmt.Printf("i value is : %v \n", i)
}
