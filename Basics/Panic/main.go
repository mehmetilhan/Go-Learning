package main

import "fmt"

func main() {
	square(5)
}

func square(v int) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error!!")
			fmt.Println("Defer worked!")
		}
	}()

	panic("------ Panic ------") // after this not work anything
	fmt.Println(v * v)
}
