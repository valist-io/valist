package git

import (
	"fmt"
	"math/big"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5/plumbing/format/pktline"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/server"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
)

type handler struct {
	client types.CoreAPI
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &handler{client}

	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}/{tag}/git-receive-pack", handler.receivePack).Methods(http.MethodPost)
	router.HandleFunc("/{org}/{repo}/git-upload-pack", handler.uploadPack).Methods(http.MethodPost)
	router.HandleFunc("/{org}/{repo}/{tag}/git-upload-pack", handler.uploadPack).Methods(http.MethodPost)
	router.HandleFunc("/{org}/{repo}/info/refs", handler.advertisedRefs).Methods(http.MethodGet)
	router.HandleFunc("/{org}/{repo}/{tag}/info/refs", handler.advertisedRefs).Methods(http.MethodGet)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *handler) advertisedRefs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	service := req.URL.Query().Get("service")

	var sess transport.Session
	var err error

	switch service {
	case transport.UploadPackServiceName:
		server := server.NewServer(&storageLoader{h.client, vars["org"], vars["repo"]})
		sess, err = server.NewUploadPackSession(nil, nil)
	case transport.ReceivePackServiceName:
		server := server.NewServer(&memLoader{})
		sess, err = server.NewReceivePackSession(nil, nil)
	default:
		http.NotFound(w, req)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refs, err := sess.AdvertisedReferences()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", fmt.Sprintf("application/x-%s-advertisement", service))
	w.Header().Add("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)

	enc := pktline.NewEncoder(w)
	enc.EncodeString(fmt.Sprintf("# service=%s\n", service)) //nolint:errcheck
	enc.Flush()                                              //nolint:errcheck

	refs.Encode(w) //nolint:errcheck
}

func (h *handler) uploadPack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	sessreq := packp.NewUploadPackRequest()
	if err := sessreq.Decode(req.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	server := server.NewServer(&storageLoader{h.client, vars["org"], vars["repo"]})
	sess, err := server.NewUploadPackSession(nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessres, err := sess.UploadPack(ctx, sessreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Type", "application/x-git-upload-pack-result")
	w.WriteHeader(http.StatusOK)

	sessres.Encode(w) //nolint:errcheck
}

func (h *handler) receivePack(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)

	if req.ContentLength == 4 {
		return // figure out why initial 0000 body is sent
	}

	res, err := h.client.ResolvePath(ctx, fmt.Sprintf("%s/%s", vars["org"], vars["repo"]))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loader := &tmpLoader{}
	server := server.NewServer(loader)

	sess, err := server.NewReceivePackSession(nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(loader.tmp)

	sessreq := packp.NewReferenceUpdateRequest()
	if err := sessreq.Decode(req.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessres, err := sess.ReceivePack(ctx, sessreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := loader.repo.Storer.PackRefs(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	releaseCID, err := h.client.Storage().WriteFile(ctx, loader.tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	release := &types.Release{
		Tag:        vars["tag"],
		ReleaseCID: releaseCID,
		MetaCID:    releaseCID,
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

	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Type", "application/x-git-receive-pack-result")
	w.WriteHeader(http.StatusOK)

	sessres.Encode(w) //nolint:errcheck
}
