package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed relay/out
//go:embed relay/out/_next
//go:embed relay/out/_next/static/chunks/pages/*.js
//go:embed relay/out/_next/static/*/*.js
var valistFS embed.FS

var httpFS http.FileSystem

func init() {
	subFS, err := fs.Sub(valistFS, "relay/out")
	if err != nil {
		panic("failed to get valist.io sub fs")
	}
	httpFS = http.FS(subFS)
}

func NewHandler() http.Handler {
	return http.FileServer(httpFS)
}
