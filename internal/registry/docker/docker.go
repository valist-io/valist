package docker

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/registry/internal/core/types"
)

type Handler struct {
	client  types.CoreAPI
	uploads map[string]int64
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &Handler{
		client:  client,
		uploads: make(map[string]int64),
	}

	router := mux.NewRouter()
	router.HandleFunc("/v2/", handler.getVersion).Methods(http.MethodGet)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/", handler.getBlob).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/", handler.postBlob).Methods(http.MethodPost)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/{uuid}", handler.patchBlob).Methods(http.MethodPatch)
	router.HandleFunc("/v2/{org}/{repo}/blobs/uploads/{uuid}", handler.putBlob).Methods(http.MethodPut)
	router.HandleFunc("/v2/{org}/{repo}/manifests/{ref}", handler.putManifest).Methods(http.MethodPut)

	// HEAD /v2/valist/docker/blobs/sha256:fd3acdcea5682abced546ec19fb6ebee725c5184e5d91614c469c0a79e67f2d0
	// DELETE /v2/<name>/blobs/uploads/<uuid>
	// POST /v2/<name>/blobs/uploads/?mount=<digest>&from=<repository name>

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *Handler) getVersion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func (h *Handler) getBlob(w http.ResponseWriter, req *http.Request) {
	// ctx := req.Context()
	// vars := mux.Vars(req)

	// orgName := vars["org"]
	// repoName := vars["repo"]

	// orgID, err := h.client.GetOrganizationID(ctx, orgName)
	// if err == types.ErrOrganizationNotExist {
	// 	http.NotFound(w, req)
	// 	return
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// repo, err := client.GetRepository(ctx, orgID, repoName)
	// if err == types.ErrRepositoryNotExist {
	// 	http.NotFound(w, req)
	// 	return
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	return
}

func (h *Handler) postBlob(w http.ResponseWriter, req *http.Request) {
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

func (h *Handler) patchBlob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	uuid := vars["uuid"]
	orgName := vars["org"]
	repoName := vars["repo"]

	// TODO verify Range header is valid and return 416 if not

	path := filepath.Join(os.TempDir(), uuid, "blob")
	blob, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer blob.Close()

	size, err := io.Copy(blob, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	start := h.uploads[uuid]
	h.uploads[uuid] += size

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Range", fmt.Sprintf("%d-%d", start, h.uploads[uuid]))
	w.Header().Set("Docker-Upload-UUID", uuid)
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) putBlob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	uuid := vars["uuid"]
	orgName := vars["org"]
	repoName := vars["repo"]

	// TODO check if multiple digest
	// TODO check if digest matches
	digest := req.URL.Query().Get("digest")

	w.Header().Set("Location", fmt.Sprintf("/v2/%s/%s/blobs/uploads/%s", orgName, repoName, uuid))
	w.Header().Set("Docker-Content-Digest", digest)

	if req.ContentLength == 0 {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	// TODO verify Range header is valid and return 416 if not

	path := filepath.Join(os.TempDir(), uuid, "blob")
	blob, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer blob.Close()

	_, err = io.Copy(blob, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) putManifest(w http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(data))
}
