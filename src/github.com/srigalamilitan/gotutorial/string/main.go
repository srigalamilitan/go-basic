package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var myString = "Hello Go"
	fmt.Println(myString)
	fmt.Println(len(myString))
	for i := 0; i < len(myString); i++ {
		fmt.Println(myString[i])
	}
	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)
	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	resultContains := strings.Contains(myString, "Go")
	fmt.Println("Ada Tidak Kata Go : ", resultContains)

	resultSplit := strings.Split(myString, " ")
	for _, v := range resultSplit {
		fmt.Println(v)
	}

	//Konversi String Menjadi Integer
	var myStringValue = "10"
	resultConvertsi, err := strconv.Atoi(myStringValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultConvertsi * 5)

	// Konversi Integer Menjadi String
	resultConversiToString := strconv.Itoa(resultConvertsi)
	fmt.Println(resultConversiToString + "test")
}
