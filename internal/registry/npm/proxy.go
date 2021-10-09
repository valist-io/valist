package npm

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/storage"
)

const DefaultRegistry = "https://registry.npmjs.org"

type proxy struct {
	client types.CoreAPI
	host   string
}

func NewProxy(client types.CoreAPI, host string) http.Handler {
	proxy := &proxy{client, host}

	router := mux.NewRouter()
	router.HandleFunc("/{name}", proxy.getMetadata).Methods(http.MethodGet)
	router.HandleFunc("/{scope}/{name}", proxy.getMetadata).Methods(http.MethodGet)
	router.HandleFunc("/-/{name}/{version}", proxy.getTarball).Methods(http.MethodGet)
	router.HandleFunc("/-/{scope}/{name}/{version}", proxy.getTarball).Methods(http.MethodGet)

	return handlers.LoggingHandler(os.Stdout, router)
}

func (p *proxy) cacheMetadata(id string) (*Metadata, error) {
	txn := p.client.Database().NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(id))
	if err != nil {
		return nil, err
	}

	data, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	var meta Metadata
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

func (p *proxy) cacheTarball(ctx context.Context, id string) (storage.File, error) {
	txn := p.client.Database().NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(id))
	if err != nil {
		return nil, err
	}

	data, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	return p.client.Storage().Open(ctx, string(data))
}

func (p *proxy) fetchMetadata(id string) (*Metadata, error) {
	cached, err := p.cacheMetadata(id)
	if err == nil {
		return cached, nil
	}

	if err != badger.ErrKeyNotFound {
		return nil, err
	}

	res, err := http.Get(DefaultRegistry + id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("failed to get npm package: status=%d body=%s", res.StatusCode, body)
	}

	var meta Metadata
	if err := json.Unmarshal(body, &meta); err != nil {
		return nil, err
	}

	txn := p.client.Database().NewTransaction(true)
	defer txn.Discard()

	entry := badger.NewEntry([]byte(id), body)
	entry = entry.WithTTL(5 * time.Minute)

	if err := txn.SetEntry(entry); err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return &meta, nil
}

func (p *proxy) fetchTarball(ctx context.Context, id string) (storage.File, error) {
	meta, err := p.cacheMetadata(path.Dir(id))
	if err != nil {
		return nil, err
	}

	cached, err := p.cacheTarball(ctx, id)
	if err == nil {
		return cached, nil
	}

	if err != badger.ErrKeyNotFound {
		return nil, err
	}

	version, ok := meta.Versions[path.Base(id)]
	if !ok {
		return nil, fmt.Errorf("invalid package version")
	}

	res, err := http.Get(version.Dist.Tarball)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("failed to get npm package: status=%d body=%s", res.StatusCode, body)
	}

	tarPaths, err := p.client.Storage().Write(ctx, body)
	if err != nil {
		return nil, err
	}

	txn := p.client.Database().NewTransaction(true)
	defer txn.Discard()

	if err := txn.Set([]byte(id), []byte(tarPaths[0])); err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return p.client.Storage().Open(ctx, tarPaths...)
}

func (p *proxy) getMetadata(w http.ResponseWriter, req *http.Request) {
	meta, err := p.fetchMetadata(req.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for semver, version := range meta.Versions {
		version.Dist.Tarball = fmt.Sprintf("http://%s/-/%s/%s", p.host, meta.Name, semver)
		meta.Versions[semver] = version
	}

	if err := json.NewEncoder(w).Encode(meta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (p *proxy) getTarball(w http.ResponseWriter, req *http.Request) {
	file, err := p.fetchTarball(req.Context(), strings.TrimPrefix(req.URL.Path, "/-"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
