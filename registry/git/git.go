package git

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/format/pktline"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/server"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist"
	"github.com/valist-io/valist/core/types"
)

const (
	GitDirName = ".git"
)

type handler struct {
	client valist.API
}

func NewHandler(client valist.API) http.Handler {
	handler := &handler{client}

	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}/git-receive-pack", handler.receivePack).Methods(http.MethodPost)
	router.HandleFunc("/{org}/{repo}/git-upload-pack", handler.uploadPack).Methods(http.MethodPost)
	router.HandleFunc("/{org}/{repo}/info/refs", handler.advertisedRefs).Methods(http.MethodGet)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *handler) error(w http.ResponseWriter, msg string, status int) {
	log.Println(msg)
	http.Error(w, msg, status)
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
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refs, err := sess.AdvertisedReferences()
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
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
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	server := server.NewServer(&storageLoader{h.client, vars["org"], vars["repo"]})
	sess, err := server.NewUploadPackSession(nil, nil)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessres, err := sess.UploadPack(ctx, sessreq)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
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
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loader := &tmpLoader{}
	server := server.NewServer(loader)

	sess, err := server.NewReceivePackSession(nil, nil)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(loader.tmp)

	sessreq := packp.NewReferenceUpdateRequest()
	if err := sessreq.Decode(req.Body); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sessres, err := sess.ReceivePack(ctx, sessreq)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := loader.repo.Storer.CountLooseRefs()
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count != 1 {
		h.error(w, "cannot push more than one ref", http.StatusInternalServerError)
		return
	}

	refs, err := loader.repo.References()
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tag, err := refs.Next()
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	head := plumbing.NewHashReference(plumbing.HEAD, tag.Hash())
	if err := loader.repo.Storer.SetReference(head); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := loader.repo.Storer.PackRefs(); err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dir, err := h.client.Storage().WriteFile(ctx, loader.tmp)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	releaseMeta := &types.ReleaseMeta{
		Name:      fmt.Sprintf("%s/%s/%s", res.OrgName, res.RepoName, tag.Name().String()),
		Artifacts: make(map[string]types.Artifact),
	}

	// TODO calculate shasum
	releaseMeta.Artifacts[GitDirName] = types.Artifact{
		Provider: dir,
	}

	releaseData, err := json.Marshal(releaseMeta)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	releasePath, err := h.client.Storage().Write(ctx, releaseData)
	if err != nil {
		h.error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	release := &types.Release{
		Tag:        tag.Name().Short(),
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

	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Type", "application/x-git-receive-pack-result")
	w.WriteHeader(http.StatusOK)

	sessres.Encode(w) //nolint:errcheck
}
