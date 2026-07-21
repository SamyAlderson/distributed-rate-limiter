package main

import (
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/coreos/etcd/etcdclient"
)

func TestRateLimiter(t *testing.T) {
	// Create a new client for the rate limiter
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	// Acquire a lock for 10 seconds
	lock, err := client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the lock is acquired
	if lock == nil {
		t.Errorf("Expected lock to be acquired, but got nil")
	}

	// Release the lock
	if err := client.ReleaseLock(context.Background(), "test_key"); err != nil {
		t.Fatal(err)
	}

	// Try to acquire the lock again while it's still held
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the lock's expiration is as expected
	expiration := lock.Expiration()
	if expiration.Add(5 * time.Second).Before(time.Now()) {
		t.Errorf("Expected lock expiration to be in the future, but got %v", expiration)
	}
}

func TestRateLimiterExpired(t *testing.T) {
	// Create a new client for the rate limiter
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	// Acquire a lock for 10 seconds and wait for it to expire
	lock, err := client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Sleep for a bit to let the lock expire
	time.Sleep(6 * time.Second)

	// Try to acquire the lock again to verify that it's expired
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the lock's expiration is as expected
	expiration := lock.Expiration()
	if expiration.Add(5 * time.Second).After(time.Now()) {
		t.Errorf("Expected lock expiration to be in the past, but got %v", expiration)
	}
}

func TestRateLimiterAcquireLockMultipleTimes(t *testing.T) {
	// Create a new client for the rate limiter
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	// Acquire a lock for 10 seconds and verify it's acquired
	lock, err := client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Try to acquire the lock again while it's still held
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the lock's expiration is as expected
	expiration := lock.Expiration()
	if expiration.Add(5 * time.Second).Before(time.Now()) {
		t.Errorf("Expected lock expiration to be in the future, but got %v", expiration)
	}

	// Release the lock and verify it's released
	if err := client.ReleaseLock(context.Background(), "test_key"); err != nil {
		t.Fatal(err)
	}

	// Try to acquire the lock again after releasing it
	_, err = client.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRateLimiterLockWithDifferentKey(t *testing.T) {
	// Create a new client for the rate limiter
	client := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	// Acquire a lock for 10 seconds with one key
	_, err := client.AcquireLock(context.Background(), "test_key1", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Try to acquire a lock for 10 seconds with a different key
	_, err = client.AcquireLock(context.Background(), "test_key2", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the lock's expiration is as expected
	expiration := client.GetLockExpiration(context.Background(), "test_key1")
	if expiration.Add(5 * time.Second).Before(time.Now()) {
		t.Errorf("Expected lock expiration to be in the future, but got %v", expiration)
	}

	// Try to acquire the lock again with the first key
	_, err = client.AcquireLock(context.Background(), "test_key1", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRateLimiterLockWithMultipleClients(t *testing.T) {
	// Create two clients for the rate limiter
	client1 := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})
	client2 := NewClient(context.Background(), &Config{
		LockTTL: 5 * time.Second,
	})

	// Acquire a lock for 10 seconds with one client
	_, err := client1.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Try to acquire a lock for 10 seconds with the other client
	_, err = client2.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the lock's expiration is as expected
	expiration := client2.GetLockExpiration(context.Background(), "test_key")
	if expiration.Add(5 * time.Second).Before(time.Now()) {
		t.Errorf("Expected lock expiration to be in the future, but got %v", expiration)
	}

	// Try to acquire the lock again with the second client
	_, err = client2.AcquireLock(context.Background(), "test_key", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}