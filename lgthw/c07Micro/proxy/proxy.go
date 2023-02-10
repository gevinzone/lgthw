package proxy

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Proxy struct {
	Client  *http.Client
	BaseUrl string
}

func NewProxy(c *http.Client, baseUrl string) *Proxy {
	return &Proxy{
		Client:  c,
		BaseUrl: baseUrl,
	}
}

func NewDefaultProxy(baseUrl string) *Proxy {
	return NewProxy(http.DefaultClient, baseUrl)
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.changeRequest(r); err != nil {
		log.Printf("error occurred during process: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error occurred during process: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p.changAndWriteResponse(w, resp)
}

func (p *Proxy) changeRequest(r *http.Request) error {
	proxyUrlRaw := p.BaseUrl + r.URL.String()
	proxyUrl, err := url.Parse(proxyUrlRaw)
	if err != nil {
		return err
	}
	r.URL = proxyUrl
	r.Host = proxyUrl.Host
	r.RequestURI = ""
	return nil
}

func (p *Proxy) copyResponse(w http.ResponseWriter, resp *http.Response) {
	var out bytes.Buffer
	_, err := out.ReadFrom(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = w.Write(out.Bytes())
}

func (p *Proxy) changAndWriteResponse(w http.ResponseWriter, r *http.Response) {
	data, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(r.StatusCode)
	_, _ = w.Write(data)
}
