package docker

type Descriptor struct {
	MediaType string `json:"mediaType"`
	Size      int64  `json:"size"`
	Digest    string `json:"digest"`
}

type Manifest struct {
	SchemaVersion int          `json:"schemaVersion"`
	MediaType     string       `json:"mediaType"`
	Config        Descriptor   `json:"config"`
	Layers        []Descriptor `json:"layers"`
}

// Digests returns a list of layer and config digests.
func (m *Manifest) Digests() []string {
	digests := []string{m.Config.Digest}
	for _, layer := range m.Layers {
		digests = append(digests, layer.Digest)
	}
	return digests
}
