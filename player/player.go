package player

import (
	"crypto/sha512"
	"hash"

	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Created             string `json:"created"`
	PasswordHash        string `json:"password"`
	digestFunc          hash.Hash
	BadPasswordAttempts int
}

func (player *Player) CreatePlayer() {
	player.digestFunc = sha512.New()
}

func (player *Player) SetPassword(password string) error {
	pBytes := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(pBytes, 14)
	player.PasswordHash = string(passwordHash)

	if err != nil {
		return err
	}
	return nil
}

func (player *Player) ComparePasswordToHash(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(player.PasswordHash), []byte(password))
	return err
}
