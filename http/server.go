package http

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/valist-io/valist"
)

func NewServer(client valist.API, addr string) (*http.Server, error) {
	// override from environment variable
	if env := os.Getenv("VALIST_API_ADDR"); env != "" {
		addr = env
	}

	handler := &handler{client}
	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}", handler.getRelease).Methods(http.MethodGet)
	router.HandleFunc("/{org}/{repo}/{tag}", handler.getRelease).Methods(http.MethodGet)

	// health check route always returns 200
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}, nil
}
