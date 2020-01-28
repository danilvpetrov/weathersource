package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danilvpetrov/weathersource/api/forecast"
	"github.com/danilvpetrov/weathersource/internal/bufferpool"
)

func (h *handlers) forecast(w http.ResponseWriter, r *http.Request) {
	f := forecast.Forecast{}
	if err := f.CopyFrom(r.Context(), h.dataAccessor); err != nil {
		http.Error(w, err.Error(), 500)
	}

	buf := bufferpool.Get()
	defer bufferpool.Put(buf)

	if err := json.NewEncoder(buf).Encode(f); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Add("Content-Type", "application/json")

	if _, err := buf.WriteTo(w); err != nil {
		panic(err)
	}
}
