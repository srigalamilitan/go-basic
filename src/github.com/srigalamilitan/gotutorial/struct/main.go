//struct type yang digunakan untuk hold segaligus dengan tipe data yang berbeda
package main

import "fmt"

type Vector struct {
	x, y int
}
type Player struct {
	ID   int
	Name string
}

func main() {
	var v Vector
	v.x = 10
	v.y = 5
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)

	player1 := Player{ID: 1, Name: "Krisna"}
	fmt.Println(player1.ID)
	fmt.Println(player1.Name)
}
