package npm

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/valist-io/registry/internal/core"
)

// Publish creates a new npm package from the given metadata.
func (h *Handler) Publish(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	rpath := strings.TrimLeft(req.URL.Path, "/@")
	parts := strings.Split(rpath, "/")
	if len(parts) != 2 {
		http.Error(w, "unscoped packages not supported", http.StatusBadRequest)
		return
	}

	orgName := parts[0]
	repoName := parts[1]

	var pack Package
	if err := json.NewDecoder(req.Body).Decode(&pack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// only support publishing latest version now
	semver, ok := pack.DistTags["latest"]
	if !ok {
		http.Error(w, "latest version required", http.StatusBadRequest)
		return
	}

	version, ok := pack.Versions[semver]
	if !ok {
		http.Error(w, "version not found", http.StatusBadRequest)
		return
	}

	attachName := fmt.Sprintf("%s-%s.tgz", pack.Name, semver)
	attach, ok := pack.Attachments[attachName]
	if !ok {
		http.Error(w, "attachment required", http.StatusBadRequest)
		return
	}

	var tarData bytes.Buffer
	buf := bytes.NewBufferString(attach.Data)
	dec := base64.NewDecoder(base64.StdEncoding, buf)

	if _, err := io.Copy(&tarData, dec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tarCID, err := h.client.WriteFile(ctx, tarData.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO calculate checksum
	version.Dist = Dist{
		Tarball: fmt.Sprintf("%s/%s", DefaultGateway, tarCID.String()),
	}

	versionData, err := json.Marshal(&version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	versionCID, err := h.client.WriteFile(ctx, versionData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println("MetaCID:", versionCID.String())
	//fmt.Println("ReleaseCID:", tarCID.String())

	orgID, err := h.client.GetOrganizationID(ctx, orgName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	release := &core.Release{
		Tag:        semver,
		ReleaseCID: tarCID,
		MetaCID:    versionCID,
	}

	vote, err := h.client.VoteRelease(ctx, &bind.TransactOpts{}, orgID, repoName, release)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		fmt.Printf("Voted to publish release %s %d/%d\n", release.Tag, vote.SigCount, vote.Threshold)
	} else {
		fmt.Printf("Approved release %s\n", release.Tag)
	}
}
