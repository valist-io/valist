package bin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist"
)

const DefaultGateway = "https://ipfs.io"

type handler struct {
	client valist.API
}

func NewHandler(client valist.API) http.Handler {
	handler := &handler{client}

	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}", handler.getRelease).Methods(http.MethodGet)
	router.HandleFunc("/{org}/{repo}/{tag}", handler.getRelease).Methods(http.MethodGet)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *handler) getRelease(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	org := vars["org"]
	repo := vars["repo"]
	tag := vars["tag"]

	if tag == "" {
		tag = "latest"
	}

	res, err := h.client.ResolvePath(ctx, fmt.Sprintf("%s/%s/%s", org, repo, tag))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	meta, err := h.client.GetReleaseMeta(ctx, res.Release.ReleaseCID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	artifactName := req.URL.Query().Get("artifact")
	artifact, ok := meta.Artifacts[artifactName]
	switch {
	case ok && artifactName != "":
		url := fmt.Sprintf("%s/%s?filename=%s", DefaultGateway, artifact.Provider, meta.Name)
		http.Redirect(w, req, url, http.StatusSeeOther)
		return
	case !ok && artifactName != "":
		http.NotFound(w, req)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(meta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
