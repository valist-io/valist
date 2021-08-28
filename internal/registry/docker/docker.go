package docker

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	cid "github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"

	"github.com/valist-io/registry/internal/core/types"
)

type Handler struct {
	client  types.CoreAPI
	blobs   map[string]cid.Cid
	uploads map[string]int64
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &Handler{
		client:  client,
		blobs:   make(map[string]cid.Cid),
		uploads: make(map[string]int64),
	}

	router := mux.NewRouter()
	router.HandleFunc("/v2/", handler.getVersion).Methods(http.MethodGet)
	router.HandleFunc("/v2/{org}/{repo}/blobs/{digest}", handler.getBlob).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/", handler.postBlob).Methods(http.MethodPost)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/{uuid}", handler.patchBlob).Methods(http.MethodPatch)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/{uuid}", handler.putBlob).Methods(http.MethodPut)
	router.HandleFunc("/v2/{org}/{repo}/manifests/{reference}", handler.putManifest).Methods(http.MethodPut)

	// DELETE /v2/<name>/blobs/uploads/<uuid>
	// POST /v2/<name>/blobs/uploads/?mount=<digest>&from=<repository name>

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *Handler) startUpload() (string, error) {
	path, err := os.MkdirTemp("", "")
	if err != nil {
		return "", err
	}

	uuid := filepath.Base(path)
	h.uploads[uuid] = 0

	return uuid, nil
}

func (h *Handler) writeUpload(uuid string, r io.Reader) (int64, error) {
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

func (h *Handler) finalizeUpload(ctx context.Context, uuid, digest string) error {
	path := filepath.Join(os.TempDir(), uuid, "blob")

	id, err := h.client.WriteFilePath(ctx, path)
	if err != nil {
		return err
	}
	defer os.RemoveAll(filepath.Dir(path))

	h.blobs[digest] = id
	delete(h.uploads, uuid)

	return nil
}

func (h *Handler) loadBlob(ctx context.Context, orgName, repoName, digest string) (files.File, error) {
	if id, ok := h.blobs[digest]; ok {
		return h.client.GetFile(ctx, id)
	}

	raw := fmt.Sprintf("%s/%s/latest/blobs/%s", orgName, repoName, digest)
	res, err := h.client.ResolvePath(ctx, raw)
	if err != nil {
		return nil, err
	}

	file, ok := res.File.(files.File)
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	return file, nil
}

func (h *Handler) getVersion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getBlob(w http.ResponseWriter, req *http.Request) {
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

	size, err := file.Size()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
	w.Header().Set("Docker-Content-Digest", digest)

	if req.Method == http.MethodHead {
		return
	}

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) postBlob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	orgName := vars["org"]
	repoName := vars["repo"]

	uuid, err := h.startUpload()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Range", "0-0")
	w.Header().Set("Docker-Upload-UUID", uuid)
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) patchBlob(w http.ResponseWriter, req *http.Request) {
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

func (h *Handler) putBlob(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	uuid := vars["uuid"]
	orgName := vars["org"]
	repoName := vars["repo"]

	// TODO verify Range header is valid and return 416 if not
	digest := req.URL.Query().Get("digest")

	if _, err := h.writeUpload(uuid, req.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.finalizeUpload(ctx, uuid, digest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Docker-Content-Digest", digest)
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) putManifest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	ref := vars["reference"]
	orgName := vars["org"]
	repoName := vars["repo"]

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manifestCID, err := h.client.WriteFile(ctx, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(manifestCID)

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/manifests/%s", orgName, repoName, ref))
	w.Header().Set("Docker-Content-Digest", fmt.Sprintf("sha256:%x", sha256.Sum256(data)))
	w.WriteHeader(http.StatusAccepted)
}
