// main.go - Main entry point for distributed rate limiter
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/etcd/etcdclient"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	// address for gRPC server
	address = "localhost:50051"
	// etcd client endpoint
	etcdEndpoint = "localhost:2379"
)

func main() {
	// setup logging
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)

	// setup etcd client
	client, err := etcdclient.New(etcdEndpoint)
	if err != nil {
		logrus.Fatal(err)
	}

	// setup gRPC server
	srv := newServer(client)
	// setup gRPC server address
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatal(err)
	}

	// start gRPC server
	go func() {
		if err := srv.Serve(lis); err != nil {
			logrus.Fatal(err)
		}
	}()

	// start gRPC server on separate goroutine
	go func() {
		if err := grpc.WaitForServerContext(context.Background(), lis.Context()); err != nil {
			logrus.Fatal(err)
		}
	}()

	// handle interrupt signal
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	// stop gRPC server
	srv.GracefulStop()

	// close etcd client
	if err := client.Close(); err != nil {
		logrus.Fatal(err)
	}
}

// newServer creates and returns a new gRPC server instance
func newServer(client *etcdclient.Client) *server {
	return &server{
		client: client,
	}
}

// server is a gRPC server serving rate limiter services
type server struct {
	client *etcdclient.Client
}

// GracefulStop stops the gRPC server gracefully
func (s *server) GracefulStop() {
	s.client.Close()
}

func init() {
	// setup etcd client endpoint
	endpoint := os.Getenv("ETCD_ENDPOINT")
	if endpoint != "" {
		etcdEndpoint = endpoint
	}
}