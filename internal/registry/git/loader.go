package git

import (
	"context"
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/valist-io/valist/internal/core/types"
)

type memLoader struct{}

func (l *memLoader) Load(ep *transport.Endpoint) (storer.Storer, error) {
	repo, err := git.Init(memory.NewStorage(), nil)
	if err != nil {
		return nil, err
	}

	return repo.Storer, nil
}

type tmpLoader struct {
	tmp  string
	repo *git.Repository
}

func (l *tmpLoader) Load(ep *transport.Endpoint) (storer.Storer, error) {
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainInit(tmp, true)
	if err != nil {
		return nil, err
	}

	l.tmp = tmp
	l.repo = repo

	return repo.Storer, nil
}

type storageLoader struct {
	client   types.CoreAPI
	orgName  string
	repoName string
}

func (l *storageLoader) Load(ep *transport.Endpoint) (storer.Storer, error) {
	ctx := context.Background()

	res, err := l.client.ResolvePath(ctx, fmt.Sprintf("%s/%s/latest", l.orgName, l.repoName))
	if err != nil {
		return nil, err
	}

	meta, err := l.client.GetReleaseMeta(ctx, res.Release.ReleaseCID)
	if err != nil {
		return nil, err
	}

	artifact, ok := meta.Artifacts[GitDirName]
	if !ok {
		return nil, fmt.Errorf("artifact not found")
	}

	storage := l.client.Storage()
	root := artifact.Provider

	fs := &storageFS{storage, root}
	cache := cache.NewObjectLRUDefault()

	return filesystem.NewStorage(fs, cache), nil
}
