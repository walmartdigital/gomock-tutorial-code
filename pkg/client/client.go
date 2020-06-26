package client

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// HTTPRequest ...
type HTTPRequest interface {
	Get(url string) (*resty.Response, error)
}

// ReadMessage ...
func ReadMessage(request HTTPRequest, animal string) string {
	resp, _ := request.Get(fmt.Sprintf("http://localhost:8080/%s", animal))
	return string(resp.Body())
}
