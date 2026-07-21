package distributed_rate_limiter

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestClient_New(t *testing.T) {
	// Create a sample configuration
	config := NewConfig()
	config.EnableLogging = false

	// Create a client with the sample configuration
	client, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
	}

	// Test that the client is not nil
	assert.NotNil(t, client)

	// Test that the client has the expected methods
	assert.NotNil(t, client.GetRateLimit)
	assert.NotNil(t, client.SetRateLimit)
}

func TestClient_GetRateLimit(t *testing.T) {
	// Create a sample client
	client, _ := NewClient(NewConfig())

	// Test that getting a rate limit returns an error
	_, err := client.GetRateLimit(context.Background(), "invalid-key")
	if err == nil {
		t.Fatal("expected error")
	}

	// Test that getting a rate limit returns the expected value
	rateLimit, err := client.GetRateLimit(context.Background(), "valid-key")
	if err != nil {
		t.Fatal(err)
	}
	if rateLimit == nil {
		t.Fatal("expected rate limit to be non-nil")
	}
	assert.NotNil(t, rateLimit)
}

func TestClient_SetRateLimit(t *testing.T) {
	// Create a sample client
	client, _ := NewClient(NewConfig())

	// Test that setting a rate limit returns an error
	err := client.SetRateLimit(context.Background(), "invalid-key", 100)
	if err == nil {
		t.Fatal("expected error")
	}

	// Test that setting a rate limit returns the expected value
	err = client.SetRateLimit(context.Background(), "valid-key", 100)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_LeaderElection(t *testing.T) {
	// Create a sample client
	client, _ := NewClient(NewConfig())

	// Test that leader election returns the expected value
	leader, err := client.LeaderElection(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if leader == nil {
		t.Fatal("expected leader to be non-nil")
	}
	assert.NotNil(t, leader)
}

func TestClient_DistributedLocking(t *testing.T) {
	// Create a sample client
	client, _ := NewClient(NewConfig())

	// Test that distributed locking returns the expected value
	locked, err := client.DistributedLocking(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if !locked {
		t.Fatal("expected lock to be held")
	}
}

func TestClient_Logs(t *testing.T) {
	// Create a sample client
	client, _ := NewClient(NewConfig())

	// Test that logging returns the expected value
	logs, err := client.Logs(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, logs)
}