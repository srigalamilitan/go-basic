package main

import (
	"fmt"
	"time"
)

func main() {
	go helloGo()
	time.Sleep(1 * time.Second)
	fmt.Println("ini fungsi Main")
}
func helloGo() {
	fmt.Println("Hello Go")
}
