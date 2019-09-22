package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("U", "", "Username")
	password := flag.String("P", "", "Username")
	flag.Parse()
	result := login(*username, *password)

	if result {
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Fail")
	}

}
func login(username string, password string) bool {
	if username == "krisna" && password == "123456" {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println("Belajar Flag")
// 	config := flag.String("C", "21212121", "test Ajahh")
// 	flag.Parse()
// 	myConfig := *config
//
// 	fmt.Println(myConfig)
// }
