package ipfs

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	merkledag "github.com/ipfs/go-merkledag"
	car "github.com/ipld/go-car"
)

// ExportCAR exports a CAR file and imports it into the valist IPFS node.
func ExportCAR(ctx context.Context, dag ipld.DAGService, id cid.Cid) error {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	// stream the car file to the http request
	go func() {
		ff, err := mw.CreateFormFile("path", id.String())
		if err != nil {
			pw.CloseWithError(err) //nolint:errcheck
			return
		}
		ses := merkledag.NewSession(ctx, dag)
		err = car.WriteCar(ctx, ses, []cid.Cid{id}, ff)
		if err != nil {
			pw.CloseWithError(err) //nolint:errcheck
			return
		}
		err = mw.Close()
		pw.CloseWithError(err) //nolint:errcheck
	}()

	req, err := http.NewRequest(http.MethodPost, "https://pin.valist.io/api/v0/dag/import", pr)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("failed to add to pin.valist.io: status=%s body=%s", res.Status, body)
	}

	return nil
}
