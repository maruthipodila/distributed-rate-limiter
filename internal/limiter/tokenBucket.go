package limiter

type TokenBucket struct {
	// Maximum number of tokens in the bucket
	Capacity int 
	// Rate at which tokens are added to the bucket
	Rate     int 
	// Tokens is the current number of tokens in the bucket
	Tokens int
	// LastRefill is the last time the bucket was refilled
	LastRefill int64
}

// NewTokenBucket creates a new TokenBucket with the given capacity and rate
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		Capacity:  capacity,
		Rate:      rate,
		Tokens:    capacity,
		LastRefill: 0,
	}
}

func (TokenBucket) Allow(key string) bool {
	// Check if the token bucket has enough tokens
	// If yes, consume a token and return true
	// If no, return false
	return true
}

