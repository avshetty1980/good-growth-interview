package application

import (
	"io"
	"net/http"
)

func findById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Write([]byte("Finding by ID: " + id))
	w.WriteHeader(http.StatusOK)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	w.Write([]byte(body))
	w.WriteHeader(http.StatusCreated)
}

func loadRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /messages/{id}", findById)
	mux.HandleFunc("POST /messages", createMessage)

	return mux
}
