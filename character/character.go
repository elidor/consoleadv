package character

import (
	"localhost.andrew/consoleapp/player"
	"localhost.andrew/consoleapp/world"
)

type Character struct {
	Name           string
	Max_Hitpoints  int
	Curr_Hitpoints int
	Size           int
	Movespeed      int
	Player         *player.Player
	Items          []Item
	CurrentRoom    *world.Room
}

type Item struct {
	Name  string
	Size  int
	Class string
	Value int
}
