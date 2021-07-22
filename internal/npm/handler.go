package npm

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/valist-io/registry/internal/core"
)

type Handler struct {
	client *core.Client
}

func NewHandler(client *core.Client) http.Handler {
	return &Handler{client}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) != 2 {
		http.Redirect(w, req, fmt.Sprintf("https://registry.npmjs.org/%s", req.URL.Path), http.StatusSeeOther)
		return
	}

	orgName := strings.TrimLeft(parts[0], "@")
	repoName := parts[1]

	_, err := h.client.GetRepository(req.Context(), orgName, repoName)
	if err == core.ErrRepositoryNotExist || err == core.ErrOrganizationNotExist {
		http.Redirect(w, req, fmt.Sprintf("https://registry.npmjs.org/%s", req.URL.Path), http.StatusSeeOther)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pack := NewPackage()
	pack.ID = req.URL.Path
	pack.Name = req.URL.Path

	iter := h.client.ListReleases(orgName, repoName, big.NewInt(1), big.NewInt(10))
	err0 := iter.ForEach(req.Context(), func(release *core.Release) {
		data, err := h.client.GetReleaseMeta(req.Context(), release.MetaCID)
		if err != nil {
			log.Printf("Failed to get release meta: %v\n", err)
		}

		var version Version
		if err := json.Unmarshal(data, &version); err != nil {
			log.Printf("Failed to parse release meta: %v\n", err)
		}

		version.ID = fmt.Sprintf("%s@%s", req.URL.Path, release.Tag)
		version.Name = req.URL.Path
		version.Version = release.Tag
		version.Dist = Dist{
			Tarball: fmt.Sprintf("https://gateway.valist.io/ipfs/%s", release.ReleaseCID.String()),
		}

		pack.Versions[release.Tag] = version
		pack.DistTags["latest"] = release.Tag
	})

	if err0 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pack)
}
