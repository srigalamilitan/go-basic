package main

import (
	"fmt"

	"github.com/srigalamilitan/gotutorial/38Package/config"
	mdl "github.com/srigalamilitan/gotutorial/38Package/model"
)

func main() {
	c1 := config.GetPostgressConnection()
	fmt.Println(c1)
	player1 := mdl.Player{ID: 1, Name: "Krisna"}
	fmt.Println(player1)
}
