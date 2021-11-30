package pinning

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

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

	// use a pipe to stream request data
	pr, pw := io.Pipe()
	// write form data to the pipe writer
	mw := multipart.NewWriter(pw)
	// pipe reader will block until content is written to pipe writer
	// this is how the data is able to be streamed in a routine
	go func() {
                 // create a form file for the multipart data
		ff, err := mw.CreateFormFile("path", strings.TrimPrefix(fpath, "/ipfs/"))
		if err != nil {
			pw.CloseWithError(err) //nolint:errcheck
			return
		}
		// copy the input data to the form file
		_, err = io.Copy(ff, car)
		if err != nil {
			pw.CloseWithError(err) //nolint:errcheck
			return
		}
		
		err = mw.Close()
		pw.CloseWithError(err) //nolint:errcheck
		// create a form file for the multipart data
		ff, err := mw.CreateFormFile("path", strings.TrimPrefix(fpath, "/ipfs/"))
		if err != nil {
			pw.CloseWithError(err) //nolint:errcheck
			return
		}
		// copy the input data to the form file
		_, err = io.Copy(ff, car)
		if err != nil {
			pw.CloseWithError(err) //nolint:errcheck
			return
		}
	}()

	defer pr.Close()

	req, err := http.NewRequest(http.MethodPost, prov.host+"/api/v0/dag/import", pr)
	req.Header.Set("Content-Type", mw.FormDataContentType())
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
		return "", fmt.Errorf("failed to add to pin.valist.io: status=%s body=%s", res.Status, body)
	}

	return fpath, nil
}
