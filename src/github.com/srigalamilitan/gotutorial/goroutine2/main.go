package main

import (
	"fmt"
)

//chanel
func main() {
	done := make(chan bool)

	go helloGo(done)
	<-done
	fmt.Println("ini fungsi Main")
}
func helloGo(done chan bool) {
	fmt.Println("Hello Go")
	done <- true
}
