package registry

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/registry/docker"
	"github.com/valist-io/valist/internal/registry/git"
	"github.com/valist-io/valist/internal/registry/npm"
)

func NewServer(client types.CoreAPI, addr string) *http.Server {
	dockerHandler := docker.NewHandler(client)
	npmHandler := npm.NewHandler(client)
	gitHandler := git.NewHandler(client)

	router := mux.NewRouter()
	router.PathPrefix("/docker/").Handler(http.StripPrefix("/docker", dockerHandler))
	router.PathPrefix("/npm/").Handler(http.StripPrefix("/npm", npmHandler))
	router.PathPrefix("/git/").Handler(http.StripPrefix("/git", gitHandler))

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
