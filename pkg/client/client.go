package client

import (
	"fmt"
)

// HTTPClient ...
type HTTPClient interface {
	Get(url string) (int, []byte, error)
}

// HTTPClientFactory ...
type HTTPClientFactory interface {
	Create() HTTPClient
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
