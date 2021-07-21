package npm

import (
	"encoding/json"
	"os"
	"testing"
)

func TestParsePackage(t *testing.T) {
	data, err := os.ReadFile("testdata/package.json")
	if err != nil {
		t.Fatalf("Failed to read package.json: %v", err)
	}

	var pack Package
	if err := json.Unmarshal(data, &pack); err != nil {
		t.Fatalf("Failed to unmarshal package: %v", err)
	}
}
