package middleware

import (
	"net/http"
	"strconv"
	"time"

	ratelimit "github.com/lucdoe/open-gateway/gateway/internal/plugins/rate-limit"
)

type RateLimitMiddleware struct {
	RateLimiter ratelimit.RateLimiter
}

func NewRateLimitMiddleware(rl ratelimit.RateLimiter) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		RateLimiter: rl,
	}
}

func (rlm *RateLimitMiddleware) setRateLimitHeaders(w http.ResponseWriter, remaining int64, window time.Duration) {
	limit := rlm.RateLimiter.GetLimit()
	w.Header().Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
	w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(remaining, 10))
	w.Header().Set("X-RateLimit-Reset", window.String())
}

func (rlm *RateLimitMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		key := req.RemoteAddr
		_, remaining, window, err := rlm.RateLimiter.RateLimit(key)
		if err != nil {
			rlm.setRateLimitHeaders(w, remaining, window)
			if err.Error() == "rate limit exceeded" {
				http.Error(w, err.Error(), http.StatusTooManyRequests)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rlm.setRateLimitHeaders(w, remaining, window)
		next.ServeHTTP(w, req)
	})
}
