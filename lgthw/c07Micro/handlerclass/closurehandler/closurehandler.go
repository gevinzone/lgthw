package closurehandler

import (
	"encoding/json"
	"net/http"
)

type ClosureHandlers struct {
	storage Storage
}

func NewClosureHandlers(storage Storage) *ClosureHandlers {
	return &ClosureHandlers{storage: storage}
}

func (c *ClosureHandlers) Load(useDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		value := "default"
		if !useDefault {
			value = c.storage.Get()
		}
		p := Payload{Value: value}
		w.WriteHeader(http.StatusOK)
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}

func (c *ClosureHandlers) Set() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		value := r.FormValue("value")
		c.storage.Put(value)
		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}

type Payload struct {
	Value string `json:"value"`
}
