package docker

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/storage"
)

type handler struct {
	client  types.CoreAPI
	blobs   map[string]string
	uploads map[string]int64
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &handler{
		client:  client,
		blobs:   make(map[string]string),
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

func (h *handler) writeUpload(uuid string, r io.Reader) (int64, error) {
	path := filepath.Join(os.TempDir(), uuid, "blob")

	blob, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer blob.Close()

	size, err := io.Copy(blob, r)
	if err != nil {
		return 0, err
	}

	h.uploads[uuid] += size
	return size, nil
}

func (h *handler) moveBlob(ctx context.Context, dir storage.Directory, digest string) error {
	p, ok := h.blobs[digest]
	if !ok {
		return fmt.Errorf("blob not found")
	}

	if err := dir.Add(ctx, digest, p); err != nil {
		return err
	}

	delete(h.blobs, digest)
	return nil
}

func (h *handler) loadBlob(ctx context.Context, orgName, repoName, digest string) (storage.File, error) {
	if p, ok := h.blobs[digest]; ok {
		return h.client.Storage().Open(ctx, p)
	}

	raw := fmt.Sprintf("%s/%s/latest/%s", orgName, repoName, digest)
	res, err := h.client.ResolvePath(ctx, raw)
	if err != nil {
		return nil, err
	}

	return res.File, nil
}

func (h *handler) loadManifest(ctx context.Context, orgName, repoName, ref string) (storage.File, error) {
	raw := fmt.Sprintf("%s/%s/latest", orgName, repoName)
	res, err := h.client.ResolvePath(ctx, raw)
	if err != nil {
		return nil, err
	}

	release, err := h.client.GetRelease(ctx, res.Organization.ID, res.Repository.Name, ref)
	if err == nil {
		return h.client.Storage().Open(ctx, release.MetaCID)
	}

	return h.client.Storage().Open(ctx, fmt.Sprintf("%s/%s", res.Release.ReleaseCID, ref))
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
	w.Header().Set("Docker-Content-Digest", digest)

	if req.Method == http.MethodHead {
		return
	}

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) postBlob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	orgName := vars["org"]
	repoName := vars["repo"]

	path, err := os.MkdirTemp("", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	// TODO verify Range header is valid and return 416 if not
	start := h.uploads[uuid]

	size, err := h.writeUpload(uuid, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Range", fmt.Sprintf("%d-%d", start, start+size))
	w.Header().Set("Docker-Upload-UUID", uuid)
	w.WriteHeader(http.StatusAccepted)
}

func (h *handler) putBlob(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	uuid := vars["uuid"]
	orgName := vars["org"]
	repoName := vars["repo"]

	// TODO verify Range header is valid and return 416 if not
	digest := req.URL.Query().Get("digest")
	path := filepath.Join(os.TempDir(), uuid, "blob")

	_, err := h.writeUpload(uuid, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p, err := h.client.Storage().WriteFile(ctx, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(filepath.Dir(path))

	h.blobs[digest] = p
	delete(h.uploads, uuid)

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Docker-Content-Digest", digest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO use tee reader
	digest := fmt.Sprintf("sha256:%x", sha256.Sum256(data))

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manifestCID, err := h.client.Storage().Write(ctx, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blobs := h.client.Storage().Mkdir()
	if err := h.moveBlob(ctx, blobs, manifest.Config.Digest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := blobs.Add(ctx, digest, manifestCID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, layer := range manifest.Layers {
		if err := h.moveBlob(ctx, blobs, layer.Digest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	release := &types.Release{
		Tag:        ref,
		ReleaseCID: blobs.Path(),
		MetaCID:    manifestCID,
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	info, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO write to hash instead of reading
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
	w.Header().Set("Content-Type", "application/vnd.docker.distribution.manifest.v2+json")
	w.Header().Set("Docker-Content-Digest", fmt.Sprintf("sha256:%x", sha256.Sum256(data)))

	if req.Method == http.MethodHead {
		return
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
