package registry

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/database/badger"
	"github.com/valist-io/valist/internal/registry/docker"
	"github.com/valist-io/valist/internal/registry/git"
	"github.com/valist-io/valist/internal/registry/npm"
)

func NewServer(client types.CoreAPI, config *config.Config) (*http.Server, error) {
	addr := os.Getenv("VALIST_API_ADDR")
	if addr == "" {
		addr = config.HTTP.ApiAddr
	}
	fmt.Println("API server running on", addr)

	// TODO move to client once RPC is implemented
	database, err := badger.NewDatabase(config.DatabasePath())
	if err != nil {
		return nil, err
	}
	npmProxy := npm.NewProxy(client, database, addr+"/proxy/npm")

	dockerHandler := docker.NewHandler(client)
	gitHandler := git.NewHandler(client)
	npmHandler := npm.NewHandler(client)

	router := mux.NewRouter()
	router.PathPrefix("/v2/").Handler(dockerHandler)
	router.PathPrefix("/api/git/").Handler(http.StripPrefix("/api/git", gitHandler))
	router.PathPrefix("/api/npm/").Handler(http.StripPrefix("/api/npm", npmHandler))
	router.PathPrefix("/proxy/npm/").Handler(http.StripPrefix("/proxy/npm", npmProxy))

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
