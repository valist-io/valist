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
	router.HandleFunc("/{org}/{repo}/{tag}", handler.getPackage).Methods(http.MethodGet)
	router.HandleFunc("/{org}/{repo}", handler.putPackage).Methods(http.MethodPut)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *Handler) writeAttachment(ctx context.Context, pack *Package, semver string) error {
	version, ok := pack.Versions[semver]
	if !ok {
		return fmt.Errorf("version not found")
	}

	attachName := fmt.Sprintf("%s-%s.tgz", pack.Name, semver)
	attach, ok := pack.Attachments[attachName]
	if !ok {
		return fmt.Errorf("attachment not found")
	}

	var tarData bytes.Buffer
	buf := bytes.NewBufferString(attach.Data)
	dec := base64.NewDecoder(base64.StdEncoding, buf)

	if _, err := io.Copy(&tarData, dec); err != nil {
		return err
	}

	tarCID, err := h.client.WriteFile(ctx, tarData.Bytes())
	if err != nil {
		return err
	}

	// TODO calculate checksum
	version.Dist = Dist{
		Tarball: fmt.Sprintf("%s/%s", DefaultGateway, tarCID.String()),
	}

	return nil
}

func (h *Handler) getPackage(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	tag := vars["tag"]
	orgName := vars["org"]
	repoName := vars["repo"]

	if tag == "" {
		tag = "latest"
	}

	res, err := h.client.ResolvePath(ctx, fmt.Sprintf("%s/%s/%s", orgName, repoName, tag))
	if err == types.ErrOrganizationNotExist {
		http.Redirect(w, req, DefaultGateway+req.URL.Path, http.StatusSeeOther)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := h.client.ReadFile(ctx, res.Release.ReleaseCID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var pack Package
	if err := json.Unmarshal(data, &pack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO is this secure?
	pack.ID = req.URL.Path
	pack.Name = req.URL.Path

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

	tag, ok := pack.DistTags["latest"]
	if !ok {
		http.Error(w, "latest tag required", http.StatusBadRequest)
		return
	}

	// TODO add all attachments to single directory
	for _, semver := range pack.DistTags {
		if err := h.writeAttachment(ctx, &pack, semver); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	pack.Attachments = nil
	packData, err := json.Marshal(&pack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	packCID, err := h.client.WriteFile(ctx, packData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	release := &types.Release{
		Tag:        tag,
		ReleaseCID: packCID,
		MetaCID:    packCID,
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
