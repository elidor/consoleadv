package player

import (
	"crypto/sha512"
	"hash"
)

type Player struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created    string `json:"created"`
	Password   string `json:"password"`
	digestFunc hash.Hash
}

func (player *Player) CreatePlayer() {
	player.digestFunc = sha512.New()
}

func (player *Player) SetPassword(password string) {
	pBytes := []byte(password)
	player.Password = string(player.digestFunc.Sum([]byte(pBytes)))
}
