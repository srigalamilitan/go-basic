package main

import "fmt"

func main() {
	var hello string = "Hello"
	var strPtr *string
	strPtr = &hello

	fmt.Println(&hello)
	fmt.Println(strPtr)

	fmt.Println(hello)
	change(hello)
	fmt.Println(hello)

	fmt.Println(hello)
	changePointer(&hello)
	fmt.Println(hello)
}

// pass by value
func change(v string) {
	v = v + " Golang"
}

// pass by reference
func changePointer(v *string) {
	*v = *v + " Golang"
}
