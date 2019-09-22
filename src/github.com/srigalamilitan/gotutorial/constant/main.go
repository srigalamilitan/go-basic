package main

import (
	"fmt"
)

const A string = "INI ADALAH KONSTANT"

func main() {
	fmt.Println(A)
	const X float32 = 3
	fmt.Println(X)
	z := 1000 / X
	fmt.Println(z)
	fmt.Printf("format %T", z)

}
