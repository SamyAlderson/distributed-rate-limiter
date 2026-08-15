package main

import (
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/coreos/etcd/etcdclient"
)

func TestRateLimiter(t *testing.T) {
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	lock, err := client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	if lock == nil {
		t.Errorf("Expected lock to be acquired, but got nil")
	}

	if err := client.ReleaseLock(context.Background(), "test_key"); err != nil {
		t.Fatal(err)
	}

	// Test acquiring same lock again
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check lock expiration
	expiration := lock.Expiration()
	if expiration.Add(5 * time.Second).Before(time.Now()) {
		t.Errorf("Expected lock expiration to be in the future, but got %v", expiration)
	}
}

func TestRateLimiterExpired(t *testing.T) {
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	lock, err := client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for lock to expire
	time.Sleep(6 * time.Second)

	// Test acquiring same lock again after expiration
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check lock expiration
	expiration := lock.Expiration()
	if expiration.Add(5 * time.Second).After(time.Now()) {
		t.Errorf("Expected lock expiration to be in the past, but got %v", expiration)
	}
}

func TestRateLimiterAcquireLockMultipleTimes(t *testing.T) {
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	lock, err := client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Test acquiring same lock multiple times
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check lock expiration
	expiration := lock.Expiration()
	if expiration.Add(5 * time.Second).Before(time.Now()) {
		t.Errorf("Expected lock expiration to be in the future, but got %v", expiration)
	}

	if err := client.ReleaseLock(context.Background(), "test_key"); err != nil {
		t.Fatal(err)
	}

	// Test acquiring same lock again after release
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRateLimiterLockWithDifferentKey(t *testing.T) {
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	// Test acquiring lock with different key
	_, err := client.AcquireLock(context.Background(), "test_key1", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Test acquiring lock with different key
	_, err = client.AcquireLock(context.Background(), "test_key2", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}