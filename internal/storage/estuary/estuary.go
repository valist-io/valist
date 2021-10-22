package estuary

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/valist-io/valist/internal/storage"
	"github.com/valist-io/valist/internal/storage/ipfs"
)

type Provider struct {
	host  string
	token string
	ipfs  *ipfs.Provider
	http  *http.Client
}

func NewProvider(host, token string, ipfs *ipfs.Provider) *Provider {
	return &Provider{
		host:  host,
		token: token,
		ipfs:  ipfs,
		http:  &http.Client{},
	}
}

func (prov *Provider) Open(ctx context.Context, fpath string) (storage.File, error) {
	return prov.ipfs.Open(ctx, fpath)
}

func (prov *Provider) ReadDir(ctx context.Context, fpath string) ([]fs.FileInfo, error) {
	return prov.ipfs.ReadDir(ctx, fpath)
}

func (prov *Provider) ReadFile(ctx context.Context, fpath string) ([]byte, error) {
	return prov.ipfs.ReadFile(ctx, fpath)
}

func (prov *Provider) WriteFile(ctx context.Context, fpath string) (string, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return "", err
	}

	return prov.Write(ctx, data)
}

func (prov *Provider) Write(ctx context.Context, data []byte) (string, error) {
	fpath, err := prov.ipfs.Write(ctx, data)
	if err != nil {
		return "", err
	}

	car, err := prov.ipfs.Export(ctx, fpath)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, prov.host+"/content/add-car", car)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+prov.token)
	res, err := prov.http.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode > 299 {
		return "", fmt.Errorf("failed to add to estuary: status=%s body=%s", res.Status, body)
	}

	return fpath, nil
}
