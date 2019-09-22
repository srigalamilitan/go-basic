package main

import "fmt"

func main() {
	defer fmt.Println("Di Eksekusi Paling Akhir")
	defer fmt.Println("Di Eksekusi Paling Akhir2")

	fmt.Println("Ini Apa")
	//
	fmt.Println("Process Normal Tanpa Keyword Defer")
	for i := 1; i <= 5; i++ {
		process(i)
	}
	fmt.Println("Process dengan Keyword Defer")
	for i := 1; i <= 5; i++ {
		defer process(i)
	}
}

func process(id int) {
	fmt.Println("Process Ke ", id)
}
