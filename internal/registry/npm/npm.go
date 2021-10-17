// Package npm implements a NodeJS package registry.
// https://github.com/npm/registry/blob/master/docs/REGISTRY-API.md
package npm

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
)

const DefaultGateway = "https://ipfs.io"

type handler struct {
	client types.CoreAPI
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &handler{client}

	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}", handler.getPackage).Methods(http.MethodGet)
	router.HandleFunc("/{org}/{repo}", handler.putPackage).Methods(http.MethodPut)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *handler) error(w http.ResponseWriter, msg string, status int) {
	log.Println(msg)
	http.Error(w, msg, status)
}

func (h *handler) getPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	tag := vars["tag"]
	if tag == "" {
		tag = "latest"
	}

	file, err := h.loadPackage(ctx, vars["org"], vars["repo"], tag)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := io.Copy(w, file); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *handler) putPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	res, err := h.client.ResolvePath(ctx, req.URL.Path)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var meta Metadata
	if err := json.NewDecoder(req.Body).Decode(&meta); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tag, ok := meta.DistTags["latest"]
	if !ok {
		h.error(w, "latest tag required", http.StatusBadRequest)
		return
	}

	versions, err := h.latestVersions(ctx, vars["org"], vars["repo"])
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	latestVersion := meta.Versions[meta.DistTags["latest"]]
	var dependencies []string
	for key, _ := range latestVersion.Dependencies {
		dependencies = append(dependencies, key)
	}

	releaseMeta := &types.ReleaseMeta{
		Name:    fmt.Sprintf("%s/%s/%s", res.OrgName, res.RepoName, tag),
		Readme:  meta.Readme,
		Version: latestVersion.Version,
		// License:      latestVersion.License,
		Dependencies: dependencies,
		Artifacts:    make(map[string]types.Artifact),
	}

	// add all new versions to the existing versions
	// each new version will be a release artifact
	for semver, version := range meta.Versions {
		attachName := fmt.Sprintf("%s-%s.tgz", meta.Name, semver)
		attach, ok := meta.Attachments[attachName]
		if !ok {
			h.error(w, "attachment not found", http.StatusBadRequest)
			return
		}

		attachPath, err := h.writeAttachment(ctx, attach.Data)
		if err != nil {
			h.error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO calculate checksum
		version.Dist = Dist{
			Tarball: fmt.Sprintf("%s/%s", DefaultGateway, attachPath),
		}

		releaseMeta.Artifacts[attachName] = types.Artifact{
			Provider: attachPath,
		}

		versions[semver] = version
	}

	meta.Attachments = nil
	meta.Versions = versions

	packData, err := json.Marshal(&meta)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	packPath, err := h.client.Storage().Write(ctx, packData)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO calculate shasum
	releaseMeta.Artifacts[MetaFileName] = types.Artifact{
		Provider: packPath,
	}

	releaseData, err := json.Marshal(releaseMeta)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	releasePath, err := h.client.Storage().Write(ctx, releaseData)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	release := &types.Release{
		Tag:        tag,
		ReleaseCID: releasePath,
		MetaCID:    types.DeprecationNotice,
	}

	vote, err := h.client.VoteRelease(ctx, res.Organization.ID, res.Repository.Name, release)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		fmt.Printf("Voted to publish release %s %d/%d\n", release.Tag, vote.SigCount, vote.Threshold)
	} else {
		fmt.Printf("Approved release %s\n", release.Tag)
	}
}
