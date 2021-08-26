package npm

import (
	"net/http"

	"github.com/valist-io/registry/internal/core"
)

const (
	DefaultGateway  = "https://ipfs.io/ipfs"
	DefaultRegistry = "https://registry.npmjs.org"
)

type Handler struct {
	client core.CoreAPI
}

func NewHandler(client core.CoreAPI) http.Handler {
	return &Handler{client}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.Read(w, req)
	case http.MethodPut:
		h.Publish(w, req)
	default:
		http.NotFound(w, req)
	}
}
