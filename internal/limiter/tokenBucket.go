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

// When request comes, check if if bucket has more than 0 tokens
// If yes, decrement the token count and return true
// If no, return false
// If the bucket is empty, check if the refill interval has passed
// If yes, refill the bucket and return true
func (bucket *TokenBucket) Allow(key string) bool {
	// Check if the refill interval has passed
	if time.Since(bucket.lastRefillTime) > bucket.refillInterval {
		// Refill the bucket
		bucket.tokensCounter = bucket.capacity
		bucket.lastRefillTime = time.Now()
	}
	// Check if there are tokens available
	if bucket.tokensCounter > 0 {
		// Decrement the token count
		bucket.tokensCounter--
		return true
	}
	return false
}

