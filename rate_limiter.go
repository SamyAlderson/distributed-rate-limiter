// Package ratelimiter provides a high-performance, distributed rate limiter for scalable systems.
package ratelimiter

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Config holds the configuration for the rate limiter.
type Config struct {
	QualityOfService string
	Quota            int
	Period           time.Duration
}

// RateLimiter is the main rate limiter struct.
type RateLimiter struct {
	mu         sync.RWMutex
	quota      int
	remaining  int
	period     time.Duration
	lastReset  time.Time
	config     Config
	logger     *logrus.Logger
	locked     bool
	lockedMux  sync.Mutex
	lockedTime time.Time
}

// NewRateLimiter returns a new rate limiter instance.
func NewRateLimiter(config Config, logger *logrus.Logger) *RateLimiter {
	return &RateLimiter{
		config:     config,
		logger:     logger,
		lastReset:  time.Now(),
		quota:      config.Quota,
		period:     config.Period,
	}
}

// Acquire attempts to acquire a rate limited resource.
func (rl *RateLimiter) Acquire() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.locked {
		rl.logger.Info("Acquire failed: already locked")
		return false
	}

	now := time.Now()
	if now.Before(rl.lastReset.Add(rl.period)) {
		// If the period has not passed, return false
		rl.logger.Info("Acquire failed: period not passed")
		return false
	}

	if rl.remaining > 0 {
		rl.remaining--
		rl.logger.Info("Acquired resource")
		return true
	}

	// If the quota is exhausted, attempt to lock the resource
	rl.locked = true
	rl.lockedMux.Lock()
	rl.lockedTime = time.Now()
	rl.logger.Info("Locked resource")
	return true
}

// Release releases a rate limited resource.
func (rl *RateLimiter) Release() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if !rl.locked {
		rl.logger.Info("Release failed: not locked")
		return
	}

	now := time.Now()
	if now.Before(rl.lockedTime.Add(rl.period)) {
		// If the period has not passed, return false
		rl.logger.Info("Release failed: period not passed")
		return
	}

	rl.locked = false
	rl.logger.Info("Released resource")
}

// Reset resets the rate limiter.
func (rl *RateLimiter) Reset() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.lastReset = time.Now()
	rl.remaining = rl.quota
	rl.logger.Info("Rate limiter reset")
}