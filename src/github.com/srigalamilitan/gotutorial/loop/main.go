package main

import (
	"fmt"
)

func main() {
	fmt.Println("	for i := 0; i <= 10; i++ // for")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Perulangan ke %d\n", i)
	}
	x := 10
	var y int

	fmt.Println("for y < x --> while")
	for y < x {
		y++
		fmt.Println(y)
	}

	fmt.Println("for /do while")
	xx := 1
	for {
		fmt.Println("Helo Go")
		xx++
		if xx == 10 {
			break
		}
	}
	xxx := 0
	for {
		xxx++
		if xxx == 5 {
			continue
		}
		fmt.Println("Nilai xxx ", xxx)
		if xxx == 10 {
			break
		}
	}

}
