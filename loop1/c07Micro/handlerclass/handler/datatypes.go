package handler

type GreetingResponse struct {
	Payload struct {
		Greeting string `json:"greeting,omitempty"`
		Name     string `json:"name,omitempty"`
		Error    string `json:"error,omitempty"`
	} `json:"payload"`
	Successful bool `json:"success"`
}

type GreetingRequest struct {
	Name     string `json:"name"`
	Greeting string `json:"greeting"`
}
