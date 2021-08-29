package git

import (
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type loader struct {
	tmp  string
	repo *git.Repository
}

func (l *loader) Load(ep *transport.Endpoint) (storer.Storer, error) {
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
