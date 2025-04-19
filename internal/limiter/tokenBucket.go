package limiter

import "time"

type TokenBucket struct {
	// Maximum number of tokens in the bucket
	capacity int 
	// refill interval in seconds
	refillInterval time.Duration	
	// Current number of tokens in the bucket
	tokensCounter int
	// LastRefill is the last time the bucket was refilled
	lastRefillTime time.Time
}

// NewTokenBucket creates a new TokenBucket with the given capacity and rate
func NewTokenBucket(capacity int, refillInterval time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		refillInterval: refillInterval,
		tokensCounter: capacity,
		lastRefillTime: time.Now(),
	}
}

func (bucket *TokenBucket) Allow(key string) bool {
	// Check if the token bucket has enough tokens
	// If yes, consume a token and return true
	// If no, return false
	return true
}

