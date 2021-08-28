// Package npm implements a NodeJS package registry.
// https://github.com/npm/registry/blob/master/docs/REGISTRY-API.md
package npm

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/registry/internal/core/types"
)

const (
	DefaultGateway  = "https://ipfs.io/ipfs"
	DefaultRegistry = "https://registry.npmjs.org"
)

type Handler struct {
	client types.CoreAPI
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &Handler{client}

	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}", handler.getPackage).Methods(http.MethodGet)
	router.HandleFunc("/{org}/{repo}", handler.putPackage).Methods(http.MethodPut)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *Handler) getPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	redirect := fmt.Sprintf("%s%s", DefaultRegistry, req.URL.Path)

	res, err := h.client.ResolvePath(ctx, req.URL.Path)
	if err != nil {
		http.Redirect(w, req, redirect, http.StatusSeeOther)
		return
	}

	pack := NewPackage()
	pack.ID = req.URL.Path
	pack.Name = req.URL.Path

	iter := h.client.ListReleases(res.Organization.ID, res.Repository.Name, big.NewInt(1), big.NewInt(10))
	err0 := iter.ForEach(ctx, func(release *types.Release) {
		data, err := h.client.ReadFile(ctx, release.MetaCID)
		if err != nil {
			log.Printf("Failed to get release meta: %v\n", err)
		}

		var version Version
		if err := json.Unmarshal(data, &version); err != nil {
			log.Printf("Failed to parse release meta: %v\n", err)
		}

		version.ID = fmt.Sprintf("%s@%s", pack.ID, release.Tag)
		version.Name = pack.Name
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

	if err := json.NewEncoder(w).Encode(pack); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) putPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	res, err := h.client.ResolvePath(ctx, req.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var pack Package
	if err := json.NewDecoder(req.Body).Decode(&pack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// only support publishing latest version now
	semver, ok := pack.DistTags["latest"]
	if !ok {
		http.Error(w, "latest version required", http.StatusBadRequest)
		return
	}

	version, ok := pack.Versions[semver]
	if !ok {
		http.Error(w, "version not found", http.StatusBadRequest)
		return
	}

	attachName := fmt.Sprintf("%s-%s.tgz", pack.Name, semver)
	attach, ok := pack.Attachments[attachName]
	if !ok {
		http.Error(w, "attachment required", http.StatusBadRequest)
		return
	}

	var tarData bytes.Buffer
	buf := bytes.NewBufferString(attach.Data)
	dec := base64.NewDecoder(base64.StdEncoding, buf)

	if _, err := io.Copy(&tarData, dec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tarCID, err := h.client.WriteFile(ctx, tarData.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO calculate checksum
	version.Dist = Dist{
		Tarball: fmt.Sprintf("%s/%s", DefaultGateway, tarCID.String()),
	}

	versionData, err := json.Marshal(&version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	versionCID, err := h.client.WriteFile(ctx, versionData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	release := &types.Release{
		Tag:        semver,
		ReleaseCID: tarCID,
		MetaCID:    versionCID,
	}

	vote, err := h.client.VoteRelease(ctx, res.Organization.ID, res.Repository.Name, release)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		fmt.Printf("Voted to publish release %s %d/%d\n", release.Tag, vote.SigCount, vote.Threshold)
	} else {
		fmt.Printf("Approved release %s\n", release.Tag)
	}
}
