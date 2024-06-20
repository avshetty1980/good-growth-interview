package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMessageServiceHandlers(t *testing.T) {

	messageStore := &mockMessageStore{}
	handler := NewMessageService(messageStore)

	t.Run("createMessage should send bad request if the message payload is invalid", func(t *testing.T) {
		payload := Message{
			Content: "",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/messages", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /messages", handler.handleCreateMessage)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("getMessage should handle by ID in path", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/messages/5099803df3f4948bd2f98391", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("GET /messages/{id}", handler.handleGetByID)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

type mockMessageStore struct{}

func (m *mockMessageStore) CreateMessage(content string) (string, error) {
	return "", nil
}

func (m *mockMessageStore) GetMessage(id string) (string, error) {
	return "", nil
}
