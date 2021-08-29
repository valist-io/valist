package git

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5/plumbing/format/pktline"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/server"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/registry/internal/core/types"
)

type handler struct {
	client types.CoreAPI
}

func NewHandler(client types.CoreAPI) http.Handler {
	handler := &handler{client}

	router := mux.NewRouter()
	router.HandleFunc("/{org}/{repo}/git-receive-pack", handler.receivePack).Methods(http.MethodPost)
	router.HandleFunc("/{org}/{repo}/info/refs", handler.advertisedRefs).Methods(http.MethodGet)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (h *handler) advertisedRefs(w http.ResponseWriter, req *http.Request) {
	service := req.URL.Query().Get("service")
	loader := &loader{}
	server := server.NewServer(loader)

	var sess transport.Session
	var err0 error

	switch service {
	case transport.UploadPackServiceName:
		sess, err0 = server.NewUploadPackSession(nil, nil)
	case transport.ReceivePackServiceName:
		sess, err0 = server.NewReceivePackSession(nil, nil)
	default:
		http.NotFound(w, req)
		return
	}

	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(loader.tmp)

	refs, err := sess.AdvertisedReferences()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", fmt.Sprintf("application/x-%s-advertisement", service))
	w.Header().Add("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)

	enc := pktline.NewEncoder(w)
	enc.EncodeString(fmt.Sprintf("# service=%s\n", service))
	enc.Flush()

	refs.Encode(w)
}

func (h *handler) receivePack(w http.ResponseWriter, req *http.Request) {
	// TODO figure out why initial 0000 body is sent
	if req.ContentLength == 4 {
		return
	}

	ctx := req.Context()
	loader := &loader{}
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

	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Type", "application/x-git-receive-pack-result")

	sessres.Encode(w)
}
