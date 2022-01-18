package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/valist-io/valist"
)

const DefaultGateway = "https://gateway.valist.io"

type handler struct {
	client valist.API
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

	artifactName := req.URL.Query().Get("artifact")
	artifact, ok := res.Release.Artifacts[artifactName]
	switch {
	case ok && artifactName != "":
		url := fmt.Sprintf("%s/%s?filename=%s", DefaultGateway, artifact.Provider, res.Release.Name)
		http.Redirect(w, req, url, http.StatusSeeOther)
		return
	case !ok && artifactName != "":
		http.NotFound(w, req)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res.Release); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
