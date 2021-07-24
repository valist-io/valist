package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/npm"
)

// Route wraps an http handler func that returns an error.
type Route func(http.ResponseWriter, *http.Request) error

func (r Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := r(w, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Server struct {
	client *core.Client
	http   *http.Server
}

func NewServer(client *core.Client, addr string) *Server {
	server := &Server{
		client: client,
	}

	npmHandler := npm.NewHandler(client)
	npmHandler = http.StripPrefix("/npm/", npmHandler)

	router := mux.NewRouter()
	router.PathPrefix("/npm/").Handler(npmHandler).Methods(http.MethodGet)
	router.Handle("/api/{org}", Route(server.GetOrganization)).Methods(http.MethodGet)
	router.Handle("/api/{org}/{repo}", Route(server.GetRepository)).Methods(http.MethodGet)
	router.Handle("/api/{org}/{repo}/releases", Route(server.ListReleases)).Methods(http.MethodGet)
	router.Handle("/api/{org}/{repo}/{tag}", Route(server.GetRelease)).Methods(http.MethodGet)

	server.http = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return server
}

// ListenAndServe starts the HTTP server.
func (s *Server) ListenAndServe() error {
	return s.http.ListenAndServe()
}

// Shutdown stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func (server *Server) GetOrganization(w http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)

	orgID, err := server.client.GetOrganizationID(req.Context(), vars["org"])
	if err != nil {
		return err
	}

	org, err := server.client.GetOrganization(req.Context(), orgID)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(org)
}

func (server *Server) GetRepository(w http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)

	orgID, err := server.client.GetOrganizationID(req.Context(), vars["org"])
	if err != nil {
		return err
	}

	repo, err := server.client.GetRepository(req.Context(), orgID, vars["repo"])
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(repo)
}

func (server *Server) ListReleases(w http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)

	page, limit, err := Paginate(req.URL.Query())
	if err != nil {
		return err
	}

	orgID, err := server.client.GetOrganizationID(req.Context(), vars["org"])
	if err != nil {
		return err
	}

	var releases []*core.Release
	iter := server.client.ListReleases(orgID, vars["repo"], page, limit)
	err0 := iter.ForEach(req.Context(), func(release *core.Release) {
		releases = append(releases, release)
	})

	if err0 != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(releases)
}

func (server *Server) GetRelease(w http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)

	orgID, err := server.client.GetOrganizationID(req.Context(), vars["org"])
	if err != nil {
		return err
	}

	release, err := server.client.GetRelease(req.Context(), orgID, vars["repo"], vars["tag"])
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(release)
}
