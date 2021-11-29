package main

import (
	"crypto/rand"
	"io"
)

type Manager struct {
	database map[string]interface{}
}

func (m *Manager) NewSessionID() string {
	b := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}

}
