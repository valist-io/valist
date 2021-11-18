package pinning

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/valist-io/valist/storage"
	"github.com/valist-io/valist/storage/ipfs"
)

type Provider struct {
	host string
	ipfs *ipfs.Provider
	http *http.Client
}

func NewProvider(host string, ipfs *ipfs.Provider) *Provider {
	return &Provider{
		host: host,
		ipfs: ipfs,
		http: &http.Client{},
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

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile("path", "file")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, car)
	if err != nil {
		fmt.Println("Error when copying file parts")
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, prov.host+"/api/v0/dag/import", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return "", err
	}

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
		return "", fmt.Errorf("failed to add to valist: status=%s body=%s", res.Status, body)
	}

	return fpath, nil
}
