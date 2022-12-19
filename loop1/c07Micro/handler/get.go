package handler

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	w.WriteHeader(http.StatusOK)
	data := map[string]string{"name": name}
	res, _ := json.Marshal(data)

	_, _ = w.Write(res)
}
