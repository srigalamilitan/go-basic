package main

import (
	"fmt"
)

//function has a value
func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(5, 5))
	hello := func(name string) string {
		return fmt.Sprintf("Hello %s", name)
	}
	fmt.Println(hello("Krisna"))
}

//normal case
// func main() {
// 	fmt.Println(add(5, 5))
// }
// func add(x, y int) int {
// 	return x + y
// }
