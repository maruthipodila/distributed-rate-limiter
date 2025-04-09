// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
