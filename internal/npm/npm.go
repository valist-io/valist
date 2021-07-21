// Package npm implements a read-only NodeJS package registry.
// https://github.com/npm/registry/blob/master/docs/REGISTRY-API.md
package npm

type Package struct {
	// the package name
	ID string `json:"_id"`
	// latest revision id
	Rev string `json:"_rev"`
	// the package name
	Name string `json:"name"`
	// description from the package.json
	Description string `json:"description"`
	// an object with at least one key, latest, representing dist-tags
	DistTags map[string]string `json:"dist-tags"`
	// a List of all Version objects for the Package
	Versions map[string]Version `json:"versions"`
	// full text of the latest version's README
	Readme string `json:"readme"`
	// an object containing a created and modified time stamp
	Time Time `json:"time"`
	// object with name, email, and or url of author as listed in package.json
	Author Author `json:"author"`
	// object with type and url of package repository as listed in package.json
	Repository Repository `json:"repository"`
}

type Version struct {
	// <name>@<version>
	ID string `json:"_id"`
	// package name
	Name string `json:"name"`
	// description as listed in package.json
	Description string `json:"description"`
	// version number
	Version string `json:"version"`
	// homepage listed in the package.json
	Homepage string `json:"homepage"`
	// object with type and url of package repository as listed in package.json
	Repository Repository `json:"repository"`
	// object with dependencies and versions as listed in package.json
	Dependencies map[string]string `json:"dependencies"`
	// object with devDependencies and versions as listed in package.json
	DevDependencies map[string]string `json:"devDependencies"`
	// object with scripts as listed in package.json
	Scripts map[string]string `json:"scripts"`
	// object with name, email, and or url of author as listed in package.json
	Author Author `json:"author"`
	// as listed in package.json
	License string `json:"license"`
	// full text of README file as pointed to in package.json
	Readme string `json:"readme"`
	// name of README file
	ReadmeFilename string `json:"readmeFilename"`
	// an object containing a shasum and tarball url, usually in the form of https://registry.npmjs.org/<name>/-/<name>-<version>.tgz
	Dist Dist `json:"dist"`
	// version of npm the package@version was published with
	NpmVersion string `json:"_npmVersion"`
	// an object containing the name and email of the npm user who published the package@version
	NpmUser Author `json:"_npmUser"`
	// an array of objects containing author objects as listed in package.json
	Maintainers []Author `json:"maintainers"`
}

type Time struct {
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Dist struct {
	Shasum  string `json:"shasum"`
	Tarball string `json:"tarball"`
}

func NewPackage() Package {
	return Package{
		DistTags: make(map[string]string),
		Versions: make(map[string]Version),
	}
}

func NewVersion() Version {
	return Version{
		Dependencies:    make(map[string]string),
		DevDependencies: make(map[string]string),
		Scripts:         make(map[string]string),
	}
}
