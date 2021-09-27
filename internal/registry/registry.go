package registry

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/registry/docker"
	"github.com/valist-io/valist/internal/registry/git"
	"github.com/valist-io/valist/internal/registry/npm"
)

func NewHandler(client types.CoreAPI) http.Handler {
	dockerHandler := docker.NewHandler(client)
	gitHandler := git.NewHandler(client)
	npmHandler := npm.NewHandler(client)

	router := mux.NewRouter()
	router.PathPrefix("/v2/").Handler(dockerHandler)
	router.PathPrefix("/api/git/").Handler(http.StripPrefix("/api/git", gitHandler))
	router.PathPrefix("/api/npm/").Handler(http.StripPrefix("/api/npm", npmHandler))

	return router
}
