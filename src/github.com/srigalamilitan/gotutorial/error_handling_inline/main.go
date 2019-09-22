package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	iStr := "5x"
	i, err := strconv.Atoi(iStr)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)
	} else {
		fmt.Println(i)
	}
	r, err := Div(10, 0)
	if err != nil {
		fmt.Println("Terjadi Error : ", err.Error())
	} else {
		fmt.Println(r)
	}

	// inline handling Error
	if x, err := Div(25, 5); err != nil {
		fmt.Println("Terjadi Error : ", err.Error())
	} else {
		fmt.Println(x)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak Bisa Membagi dengan 0")
	}
	return x / y, nil
}
