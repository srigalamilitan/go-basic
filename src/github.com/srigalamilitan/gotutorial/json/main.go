package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	player1 := Player{1, "Krisna", 3}

	p1, err := json.Marshal(player1)
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(string(p1))
	// fmt.Println(string(nil)) error dia

	//Unmarshal
	fmt.Println("UnMarshall==========================")
	p2 := []byte(`{"id":1,"age":2,"name":"gita"}`)
	var p3 Player
	err = json.Unmarshal(p2, &p3)
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(p3.ID)
	fmt.Println(p3.Name)
	fmt.Println(p3.Age)
}

type Player struct {
	ID   int    `json:id`
	Name string `json:name`
	Age  int    `json:age`
}
