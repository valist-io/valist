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

// Read serves an npm package with the given scoped name.
func (h *Handler) Read(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	redirect := fmt.Sprintf("%s%s", DefaultRegistry, req.URL.Path)

	rpath := strings.TrimLeft(req.URL.Path, "/@")
	parts := strings.Split(rpath, "/")
	if len(parts) != 2 {
		http.Redirect(w, req, redirect, http.StatusSeeOther)
		return
	}

	orgName := parts[0]
	repoName := parts[1]

	fmt.Println(orgName)
	fmt.Println(repoName)

	orgID, err := h.client.GetOrganizationID(ctx, orgName)
	if err == core.ErrOrganizationNotExist {
		http.Redirect(w, req, redirect, http.StatusSeeOther)
		return
	}

	pack := NewPackage()
	pack.ID = fmt.Sprintf("@%s/%s", orgName, repoName)
	pack.Name = fmt.Sprintf("@%s/%s", orgName, repoName)

	iter := h.client.ListReleases(orgID, repoName, big.NewInt(1), big.NewInt(10))
	err0 := iter.ForEach(ctx, func(release *core.Release) {
		data, err := h.client.ReadFile(ctx, release.MetaCID)
		if err != nil {
			log.Printf("Failed to get release meta: %v\n", err)
		}

		var version Version
		if err := json.Unmarshal(data, &version); err != nil {
			log.Printf("Failed to parse release meta: %v\n", err)
		}

		version.ID = fmt.Sprintf("@%s/%s@%s", orgName, repoName, release.Tag)
		version.Name = fmt.Sprintf("@%s/%s", orgName, repoName)
		version.Version = release.Tag
		version.Dist = Dist{
			Tarball: fmt.Sprintf("%s/%s", DefaultGateway, release.ReleaseCID.String()),
		}

		pack.Versions[release.Tag] = version
		pack.DistTags["latest"] = release.Tag
	})

	if err0 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pack)
}
