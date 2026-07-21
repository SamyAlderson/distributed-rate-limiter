// Package client provides a client interface for interacting with the rate limiter.
package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/coreos/etcd/etcdclient"
)

// Client represents a client for interacting with the rate limiter.
type Client struct {
	config *config.Config
	etcd   *etcdclient.EtcdClient
}

// NewClient returns a new instance of the client.
func NewClient(config *config.Config) (*Client, error) {
	etcd, err := etcdclient.NewEtcdClient(config.EtcdEndpoints)
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %w", err)
	}

	return &Client{
		config: config,
		etcd:   etcd,
	}, nil
}

// AcquireLock acquires a lock for the given key.
func (c *Client) AcquireLock(ctx context.Context, key string) error {
	lease := int64(c.config.LockLease)
	resp, err := c.etcd.Grant(ctx, lease)
	if err != nil {
		return fmt.Errorf("failed to grant lease: %w", err)
	}

	leaseID := resp.ID
	lockKey := fmt.Sprintf("%s/lock", key)

	putResp, err := c.etcd.Put(ctx, lockKey, "", &etcdclient.PutOptions{
		PrevExist: etcdclient.PrevExistNone,
		Lease:     leaseID,
	})
	if err != nil {
		return fmt.Errorf("failed to put lock: %w", err)
	}

	// Check if the lock was acquired successfully.
	if putResp.Header.Revision != 0 {
		return fmt.Errorf("lock already acquired by another client")
	}

	// Wait for the lock to expire.
	go func() {
		_, err := c.etcd.LeaseTime(ctx, leaseID)
		if err != nil {
			log.Printf("failed to check lease time: %v", err)
		} else {
			c.etcd.Revoke(ctx, leaseID)
		}
	}()

	return nil
}

// ReleaseLock releases the lock for the given key.
func (c *Client) ReleaseLock(ctx context.Context, key string) error {
	lockKey := fmt.Sprintf("%s/lock", key)

	resp, err := c.etcd.Delete(ctx, lockKey)
	if err != nil {
		return fmt.Errorf("failed to delete lock: %w", err)
	}

	// Check if the lock was released successfully.
	if resp.Deleted != 1 {
		return fmt.Errorf("failed to release lock")
	}

	return nil
}

// GetQuota returns the quota for the given key.
func (c *Client) GetQuota(ctx context.Context, key string) (int64, error) {
	resp, err := c.etcd.Get(ctx, key)
	if err != nil {
		return 0, fmt.Errorf("failed to get quota: %w", err)
	}

	val := resp.Kv[0].Value
	quota, err := strconv.ParseInt(string(val), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse quota value: %w", err)
	}

	return quota, nil
}

// UpdateQuota updates the quota for the given key.
func (c *Client) UpdateQuota(ctx context.Context, key string, quota int64) error {
	putResp, err := c.etcd.Put(ctx, key, strconv.FormatInt(quota, 10))
	if err != nil {
		return fmt.Errorf("failed to update quota: %w", err)
	}

	// Check if the quota was updated successfully.
	if putResp.Header.Revision != 0 {
		return fmt.Errorf("quota already exists")
	}

	return nil
}

type config struct {
	LockLease int64 `json:"lock_lease"`
	EtcdEndpoints []string `json:"etcd_endpoints"`
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}