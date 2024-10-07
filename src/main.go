package main

import (
	"fmt"
	"net/http"

	"github.com/Frontend-io/store-front/pkg/handlers"
)

const (
	port = 8080
)

func main() {
	http.HandleFunc("/", handlers.Home)

	portNumber := fmt.Sprintf(":%d", port)
	fmt.Printf("Server running on port%s 🔥", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
