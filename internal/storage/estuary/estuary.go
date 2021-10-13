package estuary

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
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

type addFileResponse struct {
	Cid       string   `json:"cid"`
	EstuaryId uint     `json:"estuaryId"`
	Providers []string `json:"providers"`
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
	var reqBody bytes.Buffer
	writer := multipart.NewWriter(&reqBody)

	part, err := writer.CreateFormFile("data", "data")
	if err != nil {
		return "", err
	}

	if _, err = part.Write(data); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, prov.host+"/content/add", &reqBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+prov.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := prov.http.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var reply addFileResponse
	if err := json.Unmarshal(resBody, &reply); err != nil {
		return "", err
	}

	if res.StatusCode > 299 {
		return "", fmt.Errorf("failed to add to estuary: status=%s body=%s", res.Status, resBody)
	}

	p := fmt.Sprintf("/ipfs/%s", reply.Cid)
	if err := prov.ipfs.Pin(ctx, p); err != nil {
		return "", err
	}

	return p, nil
}
