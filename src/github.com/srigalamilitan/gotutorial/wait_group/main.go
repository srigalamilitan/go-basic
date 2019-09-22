package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	res1 := <-process(1, 3, &wg)
	res2 := <-process(2, 1, &wg)
	res3 := <-process(3, 10, &wg)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	wg.Wait()
	// go func() {
	// 	fmt.Println("Hello From Go Routine")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// done := make(chan bool)
	//
	// go func() {
	// 	fmt.Println("Hello From Go Routine")
	// 	done <- true
	// }()
	//
	// <-done

}

func process(id int, delay time.Duration, wg *sync.WaitGroup) <-chan string {
	output := make(chan string)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * delay)
		output <- fmt.Sprintf("Process %d Done", id)

	}()
	return output

}
