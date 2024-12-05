package main

import (
	"encoding/json"
	"fmt"

	"localhost.andrew/consoleapp/player"
)

func main() {
	player := &player.Player{Name: "Bob"}
	player.CreatePlayer()
	player.SetPassword("asdf")

	jsonText, err := json.Marshal(player)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonText))
	err = player.ComparePasswordToHash("asddddf")
	if err != nil {
		player.BadPasswordAttempts++
		fmt.Println(err)
		fmt.Println(player.BadPasswordAttempts)
	}
}
