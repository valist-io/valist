package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValistConfig(t *testing.T) {

	config := ParseValistConfig()

	testConfig := &ValistConfig{
		ProjectType: "go",
		Org:         "test",
		Repo:        "binary",
		Tag:         "0.0.2",
		Meta:        "README.md",
		Build:       "make all",
		Out:         "dist",
	}

	assert.Equal(t, config, testConfig, "Test if parsed config matches expected result")
}
