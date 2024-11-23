package character

import "localhost.andrew/consoleapp/player"

type Character struct {
	Name      string
	Hitpoints int
	Size      int
	Movespeed int
	Player    *player.Player
	Items     []Item
}

type Item struct {
	Name  string
	Size  int
	Class string
	Value int
}

type Position struct {
	room *Room
}
