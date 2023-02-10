package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseHandler struct {
	ValidatePayload func(p *Payload) error
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		ValidatePayload: ValidatePayload,
	}
}

func (b *BaseHandler) Process(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var p Payload
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := b.ValidatePayload(&p); err != nil {
		switch err.(type) {
		case Verror:
			w.WriteHeader(http.StatusBadRequest)
			// pass the Verror along
			_, _ = w.Write([]byte(err.Error()))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK\n"))
}

type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
