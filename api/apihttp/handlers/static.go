package handlers

import (
	"net/http"
	"strings"

	"github.com/danilvpetrov/weathersource/assets"
	"github.com/mattetti/filebuffer"
)

func (h *handlers) static(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if _, err := assets.AssetInfo(path); err != nil {
		// Since SPA is served, if the file not found in the asset collection,
		// DO NOT issue 404. Instead serve the index file and let the client-side
		// router find the route.
		serveIndexFile(w, r)
		return
	}

	fs := http.FileServer(assets.AssetFile())
	fs.ServeHTTP(w, r)
}

func serveIndexFile(w http.ResponseWriter, r *http.Request) {
	const indexFile = "index.html"

	info, err := assets.AssetInfo(indexFile)
	if err != nil {
		// If index.html file is not found, issue a real 404.
		http.NotFoundHandler()
		return
	}

	// Serve the index file with the known 'text/html' content type to avoid
	// deduction of the content type within http.ServeContent() function.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeContent(
		w, r,
		indexFile,
		info.ModTime(),
		filebuffer.New(assets.MustAsset(indexFile)),
	)
}
