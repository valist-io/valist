package git

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/memory"
)

type tmpLoader struct {
	tmp string
}

func (l *tmpLoader) Load(ep *transport.Endpoint) (storer.Storer, error) {
	repo, err := git.PlainInit(l.tmp, true)
	if err != nil {
		return nil, err
	}

	return repo.Storer, nil
}

type memLoader struct{}

func (l *memLoader) Load(ep *transport.Endpoint) (storer.Storer, error) {
	repo, err := git.Init(memory.NewStorage(), nil)
	if err != nil {
		return nil, err
	}

	return repo.Storer, nil
}
