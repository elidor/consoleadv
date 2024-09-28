package main

import (
	"encoding/json"
	"fmt"

	"localhost.andrew/consoleapp/player"
)

func main() {
	player := &player.Player{Name: "Bob"}
	player.CreatePlayer()
	player.SetPassword("asdfjkkasdf")

	jsonText, err := json.Marshal(player)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonText))
}
