package ipfs

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	cid "github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	merkledag "github.com/ipfs/go-merkledag"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	car "github.com/ipld/go-car"
)

const (
	importURL = "https://pin.valist.io/api/v0/dag/import"
)

type Storage struct {
	ipfs coreiface.CoreAPI
}

func NewStorage(ctx context.Context, repoPath, apiAddr string, peers []string) (*Storage, error) {
	ipfs, err := newCoreAPI(ctx, repoPath, apiAddr)
	if err != nil {
		return nil, err
	}
	// sanity check that the api is active
	_, err = ipfs.Swarm().ListenAddrs(ctx)
	if err != nil {
		return nil, err
	}
	// wait until the bootstrap finishes
	bootstrap(ctx, ipfs, peers)
	return &Storage{ipfs}, nil
}

// ListFiles returns the contents of the directory at the given path.
func (s *Storage) ListFiles(ctx context.Context, fpath string) ([]string, error) {
	entryCh, err := s.ipfs.Unixfs().Ls(ctx, path.New(fpath))
	if err != nil {
		return nil, err
	}
	var entries []string
	for entry := range entryCh {
		if entry.Err != nil {
			return nil, entry.Err
		}
		entries = append(entries, entry.Name)
	}
	return entries, nil
}

// ReadFile returns the contents of the file at the given path.
func (s *Storage) ReadFile(ctx context.Context, fpath string) ([]byte, error) {
	node, err := s.ipfs.Unixfs().Get(ctx, path.New(fpath))
	if err != nil {
		return nil, err
	}
	file, ok := node.(files.File)
	if !ok {
		return nil, os.ErrNotExist
	}
	return io.ReadAll(file)
}

// WriteFile writes the contents of the file at the given path.
func (s *Storage) WriteFile(ctx context.Context, fpath string) (string, error) {
	stat, err := os.Stat(fpath)
	if err != nil {
		return "", err
	}
	node, err := files.NewSerialFile(fpath, true, stat)
	if err != nil {
		return "", err
	}
	p, err := s.ipfs.Unixfs().Add(ctx, node, options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	if err := s.exportCAR(ctx, p.Cid()); err != nil {
		return "", err
	}
	return p.String(), nil
}

// WriteBytes writes the given contents.
func (s *Storage) WriteBytes(ctx context.Context, data []byte) (string, error) {
	p, err := s.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	if err := s.exportCAR(ctx, p.Cid()); err != nil {
		return "", err
	}
	return p.String(), nil
}

// exportCAR exports a CAR and streams it to the Valist IPFS node.
func (s *Storage) exportCAR(ctx context.Context, id cid.Cid) error {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	// stream the car file to the http request
	go func() error {
		ff, err := mw.CreateFormFile("path", id.String())
		if err != nil {
			return pw.CloseWithError(err)
		}
		ses := merkledag.NewSession(ctx, s.ipfs.Dag())
		err = car.WriteCar(ctx, ses, []cid.Cid{id}, ff)
		if err != nil {
			return pw.CloseWithError(err)
		}
		return pw.CloseWithError(mw.Close())
	}()

	req, err := http.NewRequest(http.MethodPost, importURL, pr)
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
