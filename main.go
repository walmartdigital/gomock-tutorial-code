package main

import (
	"fmt"
	"net/http"

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
	fmt.Println(client.ReadMessage("monkeys"))
	fmt.Println(client.ReadMessage("dogs"))
}
