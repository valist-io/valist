// Package npm implements a NodeJS package registry.
// https://github.com/npm/registry/blob/master/docs/REGISTRY-API.md
package npm

import (
	"bytes"
	"context"
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

	"github.com/valist-io/valist/internal/core/types"
)

const (
	DefaultGateway = "https://ipfs.io"
	MetaFileName   = "doc.json"
)

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

func (h *handler) writeAttachments(ctx context.Context, release *types.ReleaseMeta, meta *Metadata) error {
	for _, semver := range meta.DistTags {
		version, ok := meta.Versions[semver]
		if !ok {
			return fmt.Errorf("version not found")
		}

		attachName := fmt.Sprintf("%s-%s.tgz", meta.Name, semver)
		attach, ok := meta.Attachments[attachName]
		if !ok {
			return fmt.Errorf("attachment not found")
		}

		var tarData bytes.Buffer
		buf := bytes.NewBufferString(attach.Data)
		dec := base64.NewDecoder(base64.StdEncoding, buf)

		if _, err := io.Copy(&tarData, dec); err != nil {
			return err
		}

		tarPaths, err := h.client.Storage().Write(ctx, tarData.Bytes())
		if err != nil {
			return err
		}

		// TODO calculate checksum
		version.Dist = Dist{
			Tarball: fmt.Sprintf("%s/%s", DefaultGateway, tarPaths[0]),
		}

		release.Artifacts[attachName] = types.Artifact{
			Providers: tarPaths,
		}

		meta.Versions[semver] = version
	}

	return nil
}

func (h *handler) getPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	tag := vars["tag"]
	orgName := vars["org"]
	repoName := vars["repo"]

	if tag == "" {
		tag = "latest"
	}

	res, err := h.client.ResolvePath(ctx, fmt.Sprintf("%s/%s/%s", orgName, repoName, tag))
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	releaseData, err := h.client.Storage().ReadFile(ctx, res.Release.ReleaseCID)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var release types.ReleaseMeta
	if err := json.Unmarshal(releaseData, &release); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	artifact, ok := release.Artifacts[MetaFileName]
	if !ok {
		h.error(w, "artifact not found", http.StatusBadRequest)
		return
	}

	file, err := h.client.Storage().Open(ctx, artifact.Providers...)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := io.Copy(w, file); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *handler) putPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

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

	releaseMeta := &types.ReleaseMeta{
		Name:      fmt.Sprintf("%s/%s/%s", res.OrgName, res.RepoName, tag),
		Artifacts: make(map[string]types.Artifact),
	}

	if err := h.writeAttachments(ctx, releaseMeta, &meta); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO include attachments from previous latest release
	// remove attachments before encoding
	meta.Attachments = nil

	packData, err := json.Marshal(&meta)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	packPaths, err := h.client.Storage().Write(ctx, packData)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO calculate shasum
	releaseMeta.Artifacts[MetaFileName] = types.Artifact{
		Providers: packPaths,
	}

	releaseData, err := json.Marshal(releaseMeta)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	releasePaths, err := h.client.Storage().Write(ctx, releaseData)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	release := &types.Release{
		Tag:        tag,
		ReleaseCID: releasePaths[0],
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
