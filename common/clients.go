package common

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"sync"
)

type Token string

func generateToken(password string) Token {
	return Token(hex.EncodeToString(crypto.SHA256.New().Sum([]byte(password))))
}

type UpdateUsers struct {
	// username -> token
	clients map[string]Token
	mutex   *sync.Mutex
}

func NewUpdateUsers() *UpdateUsers {
	return &UpdateUsers{
		clients: make(map[string]Token),
		mutex:   &sync.Mutex{},
	}
}

func (c *UpdateUsers) Login(username, password string) string {
	token := generateToken(password)

	c.mutex.Lock()
	c.clients[username] = token
	c.mutex.Unlock()
	return string(token)
}

func (c *UpdateUsers) Logout(token string) error {
	for user, tk := range c.clients {
		if tk == Token(token) {
			c.mutex.Lock()
			delete(c.clients, user)
			c.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("invalid token")
}
