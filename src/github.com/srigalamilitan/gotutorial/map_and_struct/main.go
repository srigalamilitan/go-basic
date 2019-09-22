package main

import "fmt"

func main() {
	mapPlayer := make(map[int]Player)
	mapPlayer[1] = Player{ID: 1, Name: "Krisna"}
	mapPlayer[2] = Player{ID: 2, Name: "Dira"}
	mapPlayer[3] = Player{ID: 3, Name: "Gita"}
	mapPlayer[4] = Player{ID: 4, Name: "Gandhi"}

	for _, v := range mapPlayer {
		fmt.Println(v)
	}
}

type Player struct {
	ID   int
	Name string
}
