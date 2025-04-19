package limiter

import "time"

type TokenBucket struct {
	// Maximum number of tokens in the bucket
	capacity int 
	// Rate at which tokens are added to the bucket
	refillRate int 
	// Tokens is the current number of tokens in the bucket
	tokensCounter int
	// LastRefill is the last time the bucket was refilled
	lastRefillTime time.Time
}

// NewTokenBucket creates a new TokenBucket with the given capacity and rate
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		refillRate: rate,
		tokensCounter: capacity,
		lastRefillTime: time.Now(),
	}
}

func (TokenBucket) Allow(key string) bool {
	// Check if the token bucket has enough tokens
	// If yes, consume a token and return true
	// If no, return false
	return true
}

