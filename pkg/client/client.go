package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// HTTPClient ...
type HTTPClient interface {
	Get(url string) (int, []byte, error)
}

// HTTPClientFactory ...
type HTTPClientFactory interface {
	Create() HTTPClient
}

// RestyClient ...
type RestyClient struct {
	client *resty.Client
}

// NewRestyClient ...
func NewRestyClient() *RestyClient {
	r := RestyClient{
		client: resty.New(),
	}
	return &r
}

// Get ...
func (r RestyClient) Get(url string) (int, []byte, error) {
	resp, err := r.client.R().Get(url)
	body := resp.Body()
	return resp.StatusCode(), body, err
}

// RestyClientFactory ...
type RestyClientFactory struct{}

// Create ...
func (f RestyClientFactory) Create() HTTPClient {
	r := NewRestyClient()
	return *r
}

// ZooClient ...
type ZooClient struct {
	client HTTPClient
}

// NewZooClient ...
func NewZooClient(factory HTTPClientFactory) *ZooClient {
	client := ZooClient{
		client: factory.Create(),
	}
	return &client
}

// ReadMessage ...
func (z *ZooClient) ReadMessage(animal string) string {
	_, body, _ := z.client.Get(fmt.Sprintf("http://localhost:8080/%s", animal))
	return string(body)
}
