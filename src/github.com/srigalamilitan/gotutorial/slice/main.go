package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice[0])
	for i, v := range mySlice {
		fmt.Println(i, v)
	}
	mySliceStr := []string{"krisna", "dira", "gita", "gandhi"}
	mySliceStr = append(mySliceStr, "krisna")
	for _, v := range mySliceStr {
		fmt.Println(v)
	}
}
