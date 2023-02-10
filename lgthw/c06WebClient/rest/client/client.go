package client

import (
	"fmt"
	"io"
	"net/http"
)

type ApiClient struct {
	*http.Client
}

func NewApiClient() *ApiClient {
	return &ApiClient{
		Client: &http.Client{},
	}
}

func (a *ApiClient) Get(url string) (int, error) {
	// 直接使用Client.Get方法
	//resp, err := a.Client.Get(url)
	//return a.returnResp(resp, err)
	return a.do(http.MethodGet, url, "", nil)
}

func (a *ApiClient) Post(url, contentType string, body io.Reader) (int, error) {
	// 直接使用Client.Post方法
	//resp, err := a.Client.Post(url, contentType, body)
	//a.printResp(resp)
	//return a.returnResp(resp, err)
	return a.do(http.MethodPost, url, contentType, body)
}

func (a *ApiClient) do(method, url, contentType string, body io.Reader) (int, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	resp, err := a.Do(req)
	a.printResp(resp)
	return a.returnResp(resp, err)
}

func (a *ApiClient) Put(url, contentType string, body io.Reader) (int, error) {
	return a.do(http.MethodPut, url, contentType, body)
}

func (a *ApiClient) Delete(url string) (int, error) {
	return a.do(http.MethodDelete, url, "", nil)
}

func (a *ApiClient) printResp(resp *http.Response) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(string(data))
}

func (a *ApiClient) returnResp(resp *http.Response, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
