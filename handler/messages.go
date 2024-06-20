package handler

import (
	"encoding/json"

	"io"
	"net/http"

	"github.com/avshetty1980/good-growth-interview/database"
)

type Message struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content"`
}

type MessageService struct {
	store database.Store
}

func NewMessageService(s database.Store) *MessageService {
	return &MessageService{
		store: s,
	}
}

func (m *MessageService) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /messages/{id}", m.handleGetByID)
	r.HandleFunc("POST /messages/", m.handleCreateMessage)
}

func (m *MessageService) handleCreateMessage(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	message := string(body)

	defer r.Body.Close()

	if message == "" {
		http.Error(w, "Message string cannot be empty", http.StatusBadRequest)
		return
	}

	id, err := m.store.CreateMessage(message)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := Message{ID: id}
	jsonResponse, err := json.Marshal(response.ID)
	if err != nil {
		http.Error(w, "Unable to create response with message id", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func (m *MessageService) handleGetByID(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	message, err := m.store.GetMessage(id)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(message))

}
