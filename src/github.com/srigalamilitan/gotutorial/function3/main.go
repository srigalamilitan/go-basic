package main

import (
	"fmt"
)

//clousureeeeeeee / fungsi sebagai return
func main() {
	nextValue := genNumber()
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	lv := love("Krisna Putra")
	fmt.Println(lv("Golang "))
	fmt.Println(lv("Java "))
}
func genNumber() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}
func love(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprintf("%s love %s", name, things)
	}
}
