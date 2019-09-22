package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Worker Pool")
	start := time.Now()
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for i := 1; i <= 5; i++ {
		go consumer(i, jobs, results)
	}

	go producer(jobs, 10)
	for i := 1; i <= 10; i++ {
		res := <-results
		fmt.Println("Hasil ke ", res)
	}
	elapse := time.Since(start)
	fmt.Println("Waktu yang dibutuhkan ", elapse)
}
func FakeHTTPRequest(x int) int {
	return x * 10
}
func producer(jobs chan<- int, size int) {
	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}
func consumer(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("Consumer ke ", id, " Mulai")
		time.Sleep(time.Second * 2)
		result <- FakeHTTPRequest(job)
	}

}
