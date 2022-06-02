package circuitBreaker

import (
	"time"

	"github.com/sony/gobreaker"
)

var (
	// Circuit breaker
	CircuitBreakerSettings = gobreaker.Settings{
		MaxRequests: uint32(1),
		Timeout:     time.Minute,
		Interval:    time.Hour,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.4 // At least 40% of requests failed
		},
	}
)

type GoCircuitBreaker interface {
	State() gobreaker.State
	Execute(f func() (interface{}, error)) (interface{}, error)
	Name() string
}

type CircuitBreaker struct {
	cBreakers map[string]GoCircuitBreaker
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		cBreakers: make(map[string]GoCircuitBreaker),
	}
}

// Execute returns a boolean indicating if the function was executed and it's return values
func (c *CircuitBreaker) Execute(key string, f func() (interface{}, error)) (bool, interface{}, error) {
	breaker, exists := c.cBreakers[key]
	if !exists {
		breaker = newCircuitBreaker(key)
		c.cBreakers[key] = breaker
	}
	res, err := breaker.Execute(f)
	if err == gobreaker.ErrOpenState || err == gobreaker.ErrTooManyRequests {
		return false, nil, nil
	}
	return true, res, err
}

func newCircuitBreaker(subscriber string) GoCircuitBreaker {
	settings := CircuitBreakerSettings
	settings.Name = "cb_" + subscriber

	return gobreaker.NewCircuitBreaker(settings)
}
