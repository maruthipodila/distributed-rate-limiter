package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password
		DB:       0,  // Default DB
		Protocol: 2,  // RESP2
	})

	ctx := context.Background()
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatalf("Redis SET failed: %v", err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatalf("Redis GET failed: %v", err)
	}
	fmt.Println("foo:", val)

	// Start HTTP server
	log.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
