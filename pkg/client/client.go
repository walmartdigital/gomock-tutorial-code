package client

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// ReadMessage ...
func ReadMessage(animal string) string {
	client := resty.New()
	resp, _ := client.R().Get(fmt.Sprintf("http://localhost:8080/%s", animal))
	return string(resp.Body())
}
