package handler

import (
	"encoding/json"
	"net/http"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var gr GreetingResponse
	if err := r.ParseForm(); err != nil {
		gr.Payload.Error = "bad request"
		if payload, err := json.Marshal(gr); err == nil {
			w.Write(payload)
			return
		}
	}
	name := r.FormValue("name")
	greeting := r.FormValue("greeting")

	w.WriteHeader(http.StatusOK)
	gr.Successful = true
	gr.Payload.Name = name
	gr.Payload.Greeting = greeting
	if payload, err := json.Marshal(gr); err == nil {
		w.Write(payload)
		return
	}
	w.Write([]byte("default"))
}
