package docs

import (
	"html/template"
	"net/http"
)

// NewHandler creates a new HTTP handler with routing.
func NewHandler() *Handler {
	router := http.NewServeMux()

	handler := &Handler{router: router}
	handler.router.HandleFunc("GET /", handler.Elements)
	handler.router.HandleFunc("GET /docs", handler.OpenAPI)

	return handler
}

// Handler provides API compatible with HTTP and REST standards.
type Handler struct {
	router *http.ServeMux
}

// ServeHTTP is used for joining handlers to HTTP server.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// Elements serves elements ui.
func (h *Handler) Elements(w http.ResponseWriter, _ *http.Request) {
	template.Must(template.New("elements").Parse(elements)).Execute(w, "./docs")
}

// Elements serves specification in OpenAPI standard.
func (h *Handler) OpenAPI(w http.ResponseWriter, _ *http.Request) {
	template.Must(template.New("docs").Parse(docs)).Execute(w, nil)
}
