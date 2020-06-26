package main

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/walmartdigital/gomock-tutorial-code/pkg/client"
)

func monkeys(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love monkeys!")
}

func dogs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love dogs!")
}

func main() {
	http.HandleFunc("/monkeys", monkeys)
	http.HandleFunc("/dogs", dogs)
	go http.ListenAndServe(":8080", nil)
	c := resty.New()
	fmt.Println(client.ReadMessage(c.R(), "monkeys"))
	fmt.Println(client.ReadMessage(c.R(), "dogs"))
}
