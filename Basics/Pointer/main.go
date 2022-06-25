package main

import "fmt"

var value = 5

func main() {

	showNormal()
	showPointer()

	println("**************")

	change(value)
	println(value)
	changeReference(&value)
	println(value)

}

func showNormal() {
	a := 2
	b := a
	a = 12
	fmt.Printf("Pointer of a : %v \n", &a)
	fmt.Printf("Pointer of b : %v \n", &b)
	fmt.Printf("Value of a : %v \n", a)
	fmt.Printf("Value of b : %v \n", b)
}

func showPointer() {
	a := 2
	b := &a
	a = 12
	fmt.Printf("Pointer of a : %v \n", &a)
	fmt.Printf("Pointer of b : %v \n", b)
	fmt.Printf("Value of a : %v \n", a)
	fmt.Printf("Value of b: %v \n", *b)
}

func change(val int) {
	val = 10
	fmt.Println("Value is : ", val)
}

func changeReference(val *int) {
	*val = 10
	fmt.Println("Value is : ", *val)
}
