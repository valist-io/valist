package core

import (
	"errors"
)

var (
	ErrOrganizationNotExist = errors.New("Organization does not exist")
	ErrRepositoryNotExist   = errors.New("Repository does not exist")
	ErrReleaseNotExist      = errors.New("Release does not exist")
)

type CoreAPI interface {
	CoreContractAPI
	CoreTransactorAPI
}

type CoreContractAPI interface {
	OrganizationAPI
	RegistryAPI
	ReleaseAPI
	RepositoryAPI
	StorageAPI
}

type CoreTransactorAPI interface {
	OrganizationTransactorAPI
}
