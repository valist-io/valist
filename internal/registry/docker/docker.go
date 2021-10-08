package docker

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
)

type handler struct {
	client  types.CoreAPI
	blobs   map[string][]string
	uploads map[string]int64
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &handler{
		client:  client,
		blobs:   make(map[string][]string),
		uploads: make(map[string]int64),
	}

	router := mux.NewRouter()
	router.HandleFunc("/v2/", handler.getVersion).Methods(http.MethodGet)
	router.HandleFunc("/v2/{org}/{repo}/blobs/{digest}", handler.getBlob).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/", handler.postBlob).Methods(http.MethodPost)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/{uuid}", handler.patchBlob).Methods(http.MethodPatch)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/{uuid}", handler.putBlob).Methods(http.MethodPut)
	router.HandleFunc("/v2/{org}/{repo}/manifests/{reference}", handler.putManifest).Methods(http.MethodPut)
	router.HandleFunc("/v2/{org}/{repo}/manifests/{reference}", handler.getManifest).Methods(http.MethodGet, http.MethodHead)

	// DELETE /v2/<name>/blobs/uploads/<uuid>
	// POST /v2/<name>/blobs/uploads/?mount=<digest>&from=<repository name>

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *handler) error(w http.ResponseWriter, msg string, status int) {
	log.Println(msg)
	http.Error(w, msg, status)
}

func (h *handler) getVersion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) getBlob(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	digest := vars["digest"]
	orgName := vars["org"]
	repoName := vars["repo"]

	file, err := h.loadBlob(ctx, orgName, repoName, digest)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	info, err := file.Stat()
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
	w.Header().Set("Docker-Content-Digest", digest)

	if req.Method == http.MethodHead {
		return
	}

	if _, err := io.Copy(w, file); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) postBlob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	orgName := vars["org"]
	repoName := vars["repo"]

	path, err := os.MkdirTemp("", "")
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuid := filepath.Base(path)
	h.uploads[uuid] = 0

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Range", "0-0")
	w.Header().Set("Docker-Upload-UUID", uuid)
	w.WriteHeader(http.StatusAccepted)
}

func (h *handler) patchBlob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	uuid := vars["uuid"]
	orgName := vars["org"]
	repoName := vars["repo"]

	start := h.uploads[uuid]
	if err := h.writeBlob(uuid, req.Body); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Docker-Upload-UUID", uuid)
	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Range", fmt.Sprintf("%d-%d", start, h.uploads[uuid]))
	w.WriteHeader(http.StatusAccepted)
}

func (h *handler) putBlob(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	uuid := vars["uuid"]
	orgName := vars["org"]
	repoName := vars["repo"]

	// TODO calculate and compare digest
	digest := req.URL.Query().Get("digest")
	path := filepath.Join(os.TempDir(), uuid, "blob")

	if err := h.writeBlob(uuid, req.Body); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paths, err := h.client.Storage().WriteFile(ctx, path)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(filepath.Dir(path))

	h.blobs[digest] = paths
	delete(h.uploads, uuid)

	w.Header().Set("Docker-Content-Digest", digest)
	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.WriteHeader(http.StatusAccepted)
}

func (h *handler) putManifest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	ref := vars["reference"]
	orgName := vars["org"]
	repoName := vars["repo"]

	res, err := h.client.ResolvePath(ctx, fmt.Sprintf("%s/%s", orgName, repoName))
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO use tee reader
	shasum := fmt.Sprintf("%x", sha256.Sum256(data))
	digest := fmt.Sprintf("sha256:%s", shasum)

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manifestPaths, err := h.client.Storage().Write(ctx, data)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	releaseMeta := &types.ReleaseMeta{
		Name:      fmt.Sprintf("%s/%s/%s", orgName, repoName, ref),
		Artifacts: make(map[string]types.Artifact),
	}

	releaseArtifact := types.Artifact{
		SHA256:    shasum,
		Providers: manifestPaths,
	}

	// add manifest to artifacts with ref and digest
	releaseMeta.Artifacts[ref] = releaseArtifact
	releaseMeta.Artifacts[digest] = releaseArtifact

	// add layers and config to release artifacts
	for _, digest := range manifest.Digests() {
		paths, err := h.findBlob(ctx, orgName, repoName, digest)
		if err != nil {
			h.error(w, err.Error(), http.StatusBadRequest)
			return
		}

		releaseMeta.Artifacts[digest] = types.Artifact{
			SHA256:    digest,
			Providers: paths,
		}

		delete(h.blobs, digest)
	}

	releaseData, err := json.Marshal(releaseMeta)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(releaseData))

	releasePaths, err := h.client.Storage().Write(ctx, releaseData)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	release := &types.Release{
		Tag:        ref,
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

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/manifests/%s", orgName, repoName, ref))
	w.Header().Set("Docker-Content-Digest", digest)
	w.WriteHeader(http.StatusAccepted)
}

func (h *handler) getManifest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	ref := vars["reference"]
	orgName := vars["org"]
	repoName := vars["repo"]

	file, err := h.loadManifest(ctx, orgName, repoName, ref)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	info, err := file.Stat()
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO write to hash instead of reading
	data, err := io.ReadAll(file)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
	w.Header().Set("Content-Type", "application/vnd.docker.distribution.manifest.v2+json")
	w.Header().Set("Docker-Content-Digest", fmt.Sprintf("sha256:%x", sha256.Sum256(data)))

	if req.Method == http.MethodHead {
		return
	}

	if _, err := w.Write(data); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
	}
}
