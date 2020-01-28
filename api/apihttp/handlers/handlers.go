package handlers

import (
	"net/http"

	"github.com/danilvpetrov/weathersource/storage"
)

type handlers struct {
	dataAccessor storage.DataAccessor
	mux          *http.ServeMux
}

// NewHandler returns an implementation of http.Handler interface.
func NewHandler(da storage.DataAccessor) http.Handler {
	hh := &handlers{
		dataAccessor: da,
		mux:          http.NewServeMux(),
	}

	hh.mux.HandleFunc("/api/forecast", hh.forecast)
	hh.mux.HandleFunc("/", hh.static)

	return hh
}

func (h *handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
