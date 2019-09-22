package main

import "fmt"

func main() {
	var players []Player
	players = []Player{Player{ID: 1, Name: "Krisna Putra"}, Player{ID: 2, Name: "Dira"}}
	players = append(players, Player{ID: 3, Name: "Nagita"})
	for _, v := range players {
		fmt.Println(v)
	}
}

type Player struct {
	ID   int
	Name string
}
