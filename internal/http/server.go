package http

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/npm"
)

type Server struct {
	client core.CoreAPI
}

func NewServer(client core.CoreAPI, addr string) *http.Server {
	server := Server{
		client: client,
	}

	router := gin.Default()
	router.GET("/api/:org", server.GetOrganization)
	router.GET("/api/:org/:repo", server.GetRepository)
	router.GET("/api/:org/:repo/releases", server.ListReleases)
	router.GET("/api/:org/:repo/:tag", server.GetRelease)
	router.GET("/npm/:org", server.GetNodePackage)
	router.GET("/npm/:org/*repo", server.GetNodePackage)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func (s Server) GetOrganization(c *gin.Context) {
	ctx := c.Request.Context()
	orgName := c.Param("org")

	orgID, err := s.client.GetOrganizationID(ctx, orgName)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	org, err := s.client.GetOrganization(ctx, orgID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, org)
}

func (s Server) GetRepository(c *gin.Context) {
	ctx := c.Request.Context()
	orgName := c.Param("org")
	repoName := c.Param("repo")

	orgID, err := s.client.GetOrganizationID(ctx, orgName)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	repo, err := s.client.GetRepository(ctx, orgID, repoName)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, repo)
}

func (s Server) ListReleases(c *gin.Context) {
	ctx := c.Request.Context()
	orgName := c.Param("org")
	repoName := c.Param("repo")

	var limit *big.Int
	if _, ok := limit.SetString(c.Query("limit"), 10); !ok {
		limit.SetInt64(10)
	}

	var page *big.Int
	if _, ok := page.SetString(c.Query("page"), 10); !ok {
		page.SetInt64(1)
	}

	orgID, err := s.client.GetOrganizationID(ctx, orgName)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var releases []*core.Release
	iter := s.client.ListReleases(orgID, repoName, page, limit)
	err0 := iter.ForEach(ctx, func(release *core.Release) {
		releases = append(releases, release)
	})

	if err0 != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, releases)
}

func (s Server) GetRelease(c *gin.Context) {
	ctx := c.Request.Context()
	orgName := c.Param("org")
	repoName := c.Param("repo")
	releaseTag := c.Param("tag")

	orgID, err := s.client.GetOrganizationID(ctx, orgName)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	release, err := s.client.GetRelease(ctx, orgID, repoName, releaseTag)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, release)
}

func (s Server) GetNodePackage(c *gin.Context) {
	ctx := c.Request.Context()
	orgName := strings.TrimLeft(c.Param("org"), "@")
	repoName := strings.TrimLeft(c.Param("repo"), "/")

	// handle unscoped package redirects
	if repoName == "" {
		redirect := fmt.Sprintf("%s/%s", npm.DefaultRegistry, orgName)
		c.Redirect(http.StatusSeeOther, redirect)
		return
	}

	// TODO move to registry mapping
	registry := npm.NewRegistry(s.client)

	pack, err := registry.GetScopedPackage(ctx, orgName, repoName)
	if err == core.ErrOrganizationNotExist {
		redirect := fmt.Sprintf("%s/%s/%s", npm.DefaultRegistry, orgName, repoName)
		c.Redirect(http.StatusSeeOther, redirect)
		return
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, pack)
}
