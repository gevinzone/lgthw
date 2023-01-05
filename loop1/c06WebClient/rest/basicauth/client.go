package basicauth

import (
	"fmt"
	"io"
	"net/http"
)

type ApiClient struct {
	*http.Client
}

func NewApiClient(username, password string) *ApiClient {
	t := &http.Transport{}
	return &ApiClient{
		Client: &http.Client{
			Transport: &ApiTransport{
				Transport: t,
				username:  username,
				password:  password,
			},
		},
	}
}

func (a *ApiClient) Get(url string) (int, error) {
	resp, err := a.Client.Get(url)
	if err != nil {
		return 0, err
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(string(data))
	}
	return resp.StatusCode, nil
}

type ApiTransport struct {
	*http.Transport
	username string
	password string
}

func (a *ApiTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(a.username, a.password)
	return a.Transport.RoundTrip(request)
}
