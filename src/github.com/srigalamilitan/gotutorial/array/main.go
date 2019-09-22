package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(arr[3])

	arrStr := [5]string{"ega", "lanri", "prima", "rayga", "krisna"}
	fmt.Println(arrStr[4])

	arrMulti := [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(arrMulti[1][2])

}
