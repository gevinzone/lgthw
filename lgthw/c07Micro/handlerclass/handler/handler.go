package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Handlers struct {
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) Post(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) PostJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req GreetingRequest
	//defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &req)
	if err != nil {
		return
	}

	var gr GreetingResponse
	name := req.Name
	greeting := req.Greeting

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
