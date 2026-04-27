package utils

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	// visitors holds the rate limiter for each IP address.
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

// visitor represents a single IP address and its rate limiter.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// init starts a background goroutine to clean up old IP entries every minute.
func init() {
	go cleanupVisitors()
}

// getVisitor returns the rate limiter for the provided IP address.
func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		// Allow 5 requests per second, burst of 10
		limiter := rate.NewLimiter(5, 10)
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// cleanupVisitors removes stale IPs that haven't been seen in 3 minutes.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// RateLimiter is an HTTP middleware that limits requests per IP.
func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract IP from RemoteAddr or Headers (for Vercel)
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		}

		limiter := getVisitor(ip)
		if !limiter.Allow() {
			SendError(w, http.StatusTooManyRequests, "Rate limit exceeded. Please slow down.", "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
