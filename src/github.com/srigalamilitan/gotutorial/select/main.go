package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang Select Chanel")
	httpRequest1 := process(1, 3)
	httpRequest2 := process(2, 3)
	httpRequest3 := process(3, 1)
	for i := 1; i <= 3; i++ {
		select {
		case p1 := <-httpRequest1:
			fmt.Println(p1)
		case p2 := <-httpRequest2:
			fmt.Println(p2)
		case p3 := <-httpRequest3:
			fmt.Println(p3)
		}
	}
}
func process(id int, delay time.Duration) <-chan string {
	output := make(chan string)
	go func() {
		fmt.Println("Process id ", id, " Started")
		time.Sleep(time.Second * delay)
		output <- fmt.Sprintf("Process %d is Done", id)
	}()
	return output
}
