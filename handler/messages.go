package handler

import (
	"fmt"
	"net/http"
)

type Message struct {
	// ID      string
	// Content string
}

func (m *Message) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Created a message")
}

func (m *Message) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get message by ID")
}
