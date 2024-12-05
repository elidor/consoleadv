package player

import (
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	player := Player{Name: "Bob"}
	player.CreatePlayer()

	if player.Name != "Bob" {
		t.Fatal("Unable to initialize Player")
	}
}
