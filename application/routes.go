package application

import (
	"net/http"

	"github.com/avshetty1980/good-growth-interview/handler"
)

// func findById(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")

// 	w.Write([]byte("Finding by ID: " + id))

// }

// func createMessage(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer r.Body.Close()
// 	w.Write([]byte(body))

// }

func loadRoutes() http.Handler {
	messageHandler := &handler.Message{}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /messages/{id}", messageHandler.GetByID)
	mux.HandleFunc("POST /messages", messageHandler.Create)

	return mux
}
