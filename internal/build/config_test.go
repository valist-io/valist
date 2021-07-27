package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValistConfig(t *testing.T) {

	config := ParseValistConfig()

	testConfig := ValistConfig{
		Type:  "go",
		Org:   "test",
		Repo:  "binary",
		Tag:   "0.0.2",
		Meta:  "README.md",
		Build: "make all",
		Out:   "dist",
		Artifacts: map[string]string{
			"linux/amd64":  "bin/lin000x/hello-world",
			"darwin/amd64": "bin/macz/hello-world",
		},
	}

	assert.Equal(t, config, testConfig, "Test if parsed config matches expected result")
}
