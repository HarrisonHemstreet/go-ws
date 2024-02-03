package main

import (
	"fmt"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler/hello"
)

func main() {
	http.HandleFunc("/", hello.HelloWorldHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
