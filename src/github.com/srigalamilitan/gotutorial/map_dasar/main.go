package main

import "fmt"

func main() {
	var mapPerson map[int]string
	mapPerson = make(map[int]string)
	mapPerson[1] = "Krisna"
	mapPerson[2] = "Dira"
	mapPerson[4] = "Wury"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}
	wury, ok := mapPerson[4]
	if !ok {
		fmt.Println("Wury Tidak ada Di map Person")
	} else {
		fmt.Println(wury)
	}
}
