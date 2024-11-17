package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	mu        sync.Mutex
	requests  map[string]int
	resetTime map[string]time.Time
	limit     int
	duration  time.Duration
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:  make(map[string]int),
		resetTime: make(map[string]time.Time),
		limit:     limit,
		duration:  duration,
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		clientIP := c.ClientIP()
		now := time.Now()

		// Reset limit if duration has passed
		if reset, exists := rl.resetTime[clientIP]; exists && now.After(reset) {
			rl.requests[clientIP] = 0
			rl.resetTime[clientIP] = now.Add(rl.duration)
		}

		if rl.requests[clientIP] < rl.limit {
			rl.requests[clientIP]++
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate limit exceeded",
			})
		}
	}
}