package registry

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/valist-io/valist"
	"github.com/valist-io/valist/core/config"
	"github.com/valist-io/valist/database/badger"
	"github.com/valist-io/valist/registry/bin"
	"github.com/valist-io/valist/registry/docker"
	"github.com/valist-io/valist/registry/git"
	"github.com/valist-io/valist/registry/npm"
)

func NewServer(client valist.API, config *config.Config) (*http.Server, error) {
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

	binHandler := bin.NewHandler(client)
	dockerHandler := docker.NewHandler(client)
	gitHandler := git.NewHandler(client)
	npmHandler := npm.NewHandler(client)
	npmProxy := npm.NewProxy(client, database, addr+"/proxy/npm")

	router := mux.NewRouter()
	router.PathPrefix("/v2/").Handler(dockerHandler)
	router.PathPrefix("/api/bin/").Handler(http.StripPrefix("/api/bin", binHandler))
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
