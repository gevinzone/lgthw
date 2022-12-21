package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GreetingJsonHandler(w http.ResponseWriter, r *http.Request) {
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
