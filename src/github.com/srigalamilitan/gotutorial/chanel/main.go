package main

import "fmt"

func main() {

	hello := make(chan string, 2)
	hello <- "krisna"
	hello <- "putra"
	close(hello)
	// fmt.Println(<-hello)
	// fmt.Println(<-hello)

	for v := range hello {
		fmt.Println(v)
	}
}
