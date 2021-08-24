package npm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/valist-io/registry/internal/core/types"
)

// TODO replace with config
const (
	DefaultGateway  = "https://ipfs.io/ipfs"
	DefaultRegistry = "https://registry.npmjs.org"
)

type Registry struct {
	client types.CoreAPI
}

func NewRegistry(client types.CoreAPI) *Registry {
	return &Registry{
		client: client,
	}
}

// GetScopedPackage returns the package from the repository with the given organization and name.
func (r *Registry) GetScopedPackage(ctx context.Context, orgName, repoName string) (*Package, error) {
	orgID, err := r.client.GetOrganizationID(ctx, orgName)
	if err != nil {
		return nil, err
	}

	pack := NewPackage()
	pack.ID = fmt.Sprintf("@%s/%s", orgName, repoName)
	pack.Name = fmt.Sprintf("@%s/%s", orgName, repoName)

	iter := r.client.ListReleases(orgID, repoName, big.NewInt(1), big.NewInt(10))
	err0 := iter.ForEach(ctx, func(release *types.Release) {
		data, err := r.client.ReadFile(ctx, release.MetaCID)
		if err != nil {
			log.Printf("Failed to get release meta: %v\n", err)
		}

		var version Version
		if err := json.Unmarshal(data, &version); err != nil {
			log.Printf("Failed to parse release meta: %v\n", err)
		}

		version.ID = fmt.Sprintf("@%s/%s@%s", orgName, repoName, release.Tag)
		version.Name = fmt.Sprintf("@%s/%s", orgName, repoName)
		version.Version = release.Tag
		version.Dist = Dist{
			Tarball: fmt.Sprintf("%s/%s", DefaultGateway, release.ReleaseCID.String()),
		}

		pack.Versions[release.Tag] = version
		pack.DistTags["latest"] = release.Tag
	})

	if err0 != nil {
		return nil, err0
	}

	return &pack, nil
}
