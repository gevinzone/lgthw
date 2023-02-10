package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient
	// 这个与DefaultClient相同，但不会把DefaultClient改掉
	//c = &http.Client{}
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	return c
}

type NopTransport struct {
}

func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}

// 确保NopTransport 实现了RoundTripper 接口
var _ http.RoundTripper = &NopTransport{}

func DoGetOps(c *http.Client, url string) error {
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func DefaultGet(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

type MyClient struct {
	*http.Client
}

func (m *MyClient) DoGetOps(url string) error {
	//resp, err := m.Client.Get(url)
	resp, err := m.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}
