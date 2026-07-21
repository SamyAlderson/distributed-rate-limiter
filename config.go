// Package config provides configuration options for the distributed rate limiter.
package config

import (
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
)

// Config holds the configuration options for the rate limiter.
type Config struct {
	// Etcd endpoint URL.
	EtcdURL string

	// Leader election lease duration.
	LeaseDuration time.Duration

	// Distributed locking timeout.
	LockTimeout time.Duration

	// Quota settings.
	QuotaSettings *QuotaSettings
}

// QuotaSettings holds the quota settings.
type QuotaSettings struct {
	// Maximum number of requests allowed within the window.
	MaxRequests int

	// Time window for the quota.
	Window time.Duration
}

// NewConfig returns a new Config instance.
func NewConfig() *Config {
	return &Config{
		LeaseDuration: 10 * time.Second,
		LockTimeout:   30 * time.Second,
	}
}

// Validate checks that the configuration is valid.
func (c *Config) Validate() error {
	if c.EtcdURL == "" {
		return fmt.Errorf("etcd URL is required")
	}

	if c.LeaseDuration <= 0 {
		return fmt.Errorf("lease duration must be greater than zero")
	}

	if c.LockTimeout <= 0 {
		return fmt.Errorf("lock timeout must be greater than zero")
	}

	if c.QuotaSettings == nil {
		return fmt.Errorf("quota settings are required")
	}

	if c.QuotaSettings.MaxRequests <= 0 {
		return fmt.Errorf("max requests must be greater than zero")
	}

	if c.QuotaSettings.Window <= 0 {
		return fmt.Errorf("window must be greater than zero")
	}

	return nil
}

// LoadConfig loads the configuration from a file.
func LoadConfig(filename string) (*Config, error) {
	// Not implemented, but could be loaded from a file or environment variables.
	return NewConfig(), nil
}

func init() {
	// Register the default logger.
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
}