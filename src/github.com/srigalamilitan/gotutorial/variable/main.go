package main

import (
	"fmt"
)

func main() {
	var x int
	x = 10

	var y float64
	y = 5.5
	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("x memiliki tipe data %T\n", x)
	fmt.Printf("y memiliki tipe data %T\n", y)

	// dynamic type declaration
	z := "krisna putra"
	i := 50
	fmt.Println(z)
	fmt.Println(i)

	fmt.Printf("z memiliki tipe data %T\n", z)
	fmt.Printf("i memiliki tipe data %T\n", i)

	// xx := i * y
	// fmt.Println(xx)
	// fmt.Printf("xx memiliki tipe data %T\n", xx)
}
