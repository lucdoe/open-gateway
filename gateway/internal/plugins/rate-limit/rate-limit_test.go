package ratelimit_test

import (
	"errors"
	"testing"
	"time"

	ratelimit "github.com/lucdoe/open-gateway/gateway/internal/plugins/rate-limit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Increment(key string, window time.Duration) (int64, error) {
	args := m.Called(key, window)
	return args.Get(0).(int64), args.Error(1)
}

func TestRateLimitService(t *testing.T) {
	mockCache := new(MockCache)
	limit := int64(10)
	window := time.Minute
	service := ratelimit.NewRateLimitService(mockCache, limit, window)

	t.Run("Under limit", func(t *testing.T) {
		mockCache.On("Increment", "123.123.123.123", window).Return(int64(1), nil).Once()
		count, remaining, _, err := service.RateLimit("123.123.123.123")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), count)
		assert.Equal(t, limit-1, remaining)
	})

	t.Run("Rate limit exceeded", func(t *testing.T) {
		mockCache.On("Increment", "123.123.123.123", window).Return(int64(11), nil).Once()
		count, remaining, _, err := service.RateLimit("123.123.123.123")
		assert.Error(t, err)
		assert.Equal(t, "rate limit exceeded", err.Error())
		assert.Equal(t, int64(11), count)
		assert.Equal(t, int64(0), remaining)
	})

	t.Run("Increment error", func(t *testing.T) {
		mockCache.On("Increment", "123.123.123.123", window).Return(int64(0), errors.New("redis error")).Once()
		_, _, _, err := service.RateLimit("123.123.123.123")
		assert.Error(t, err)
		assert.Equal(t, "redis error", err.Error())
	})
}
