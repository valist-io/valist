//go:generate ipfs get /ipns/valist.io
package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed valist.io
//go:embed valist.io/_next
//go:embed valist.io/_next/static/chunks/pages/*.js
//go:embed valist.io/_next/static/*/*.js
var valistFS embed.FS

func NewServer(bindAddr string) *http.Server {
	rootFS, err := fs.Sub(valistFS, "valist.io")
	if err != nil {
		panic("failed to get valist.io sub fs")
	}
	
	return &http.Server{
		Addr:    bindAddr,
		Handler: http.FileServer(http.FS(rootFS)),
	}
}
