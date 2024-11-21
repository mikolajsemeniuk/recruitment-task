package index

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

// Storage defines operations on datastore.
type Storage interface {
	Find(ctx context.Context, number int) (int, error)
}

// Handler provides API compatible with REST standards.
type Handler struct {
	mux   *http.ServeMux
	store Storage
}

// NewHandler creates a new HTTP handler with routing.
func NewHandler(s Storage) *Handler {
	mux := http.NewServeMux()

	handler := &Handler{mux: mux, store: s}

	handler.mux.HandleFunc("GET /{value}", handler.Find)
	// Add more endpoints below...

	return handler
}

// ServeHTTP is used for joining handlers to HTTP server.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// Find serves found index or error.
func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	input, err := NewFindInput(r)
	if err != nil {
		slog.Error("Failed to create input", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	index, err := h.store.Find(r.Context(), input.Number)
	if errors.Is(err, ErrIndexNotFound) {
		// I decided to not put slog everywhere and here it's why.
		// Logs are usually the most expensive thing because of the log pollution (from which also hard to find something) done by developers.
		// It's very common that record is not found and does not mean that application has degraded so logging this everywhere will increase costs drastically.
		// If we would like to monitor condition of our application based on status not found metric it would be better to use prometheus to avoid too high costs.
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("Failed to find index", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output := FindOutput{Index: index}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
