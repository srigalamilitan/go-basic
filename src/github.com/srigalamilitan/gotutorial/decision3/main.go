package main

import (
	"fmt"
)

func main() {
	x := 100
	switch x {
	case 60:
		fmt.Println("NILAI C")
	case 80:
		fmt.Println("NILAI B")
	case 100:
		fmt.Println("NILAI A")
	default:
		fmt.Println("NILAI TIDAK DI KETAHUI")

	}

	switch {
	case x == 60:
		fmt.Println("NILAI C")
	case x == 80:
		fmt.Println("NILAI B")
	case x == 100:
		fmt.Println("NILAI A")
	default:
		fmt.Println("NILAI TIDAK DI KETAHUI")
	}

	// type switch
	var y interface{}
	y = "Krisna Putra"
	switch y.(type) {
	case int:
		fmt.Println("Y Bertipe Data Integer")
	case string:
		fmt.Println("Y Bertipe Data String")
	case float64:
		fmt.Println("Y Bertipe Data Float6")
	default:
		fmt.Println("Tidak Di ketahui Tipe Datanya")
	}
}
