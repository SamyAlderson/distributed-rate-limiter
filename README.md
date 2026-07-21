distributed-rate-limiter
=====================

A high-performance, distributed rate limiter for scalable systems.

### What

This project provides a high-performance, distributed rate limiter for scalable systems. It uses leader election and distributed locking to ensure consistent and accurate rate limiting.

### Why

Rate limiting is a critical component of many scalable systems. It prevents abuse and ensures fair access to resources. This project provides a reliable and high-performance solution for rate limiting in distributed systems.

### Install

To install and run this project, follow these steps:

1. Clone the repository using `git clone https://github.com/samy/distributed-rate-limiter.git`
2. Navigate to the project directory using `cd distributed-rate-limiter`
3. Install dependencies using `go get`
4. Run the program using `go run main.go`

### Usage

To use the rate limiter, create a client instance and configure it with the desired quota and timeout. The client will automatically connect to the rate limiter and begin enforcing the quota.

```go
import "github.com/samy/distributed-rate-limiter/client"

// Create a new client instance
client := client.NewClient("localhost:2379", 100, 1*time.Second)

// Increment the rate limiter
err := client.Increment()
if err != nil {
    log.Fatal(err)
}
```

### Build from Source

To build the project from source, run the following commands:

1. `go build main.go`
2. `go build rate_limiter.go`
3. `go build client.go`
4. `go build config.go`

### Project Structure

The project is structured as follows:

* `go.mod` and `go.sum` contain the Go module declaration and checksums
* `main.go` is the main entry point of the program
* `rate_limiter.go` contains the rate limiter implementation
* `client.go` contains the client for interacting with the rate limiter
* `config.go` contains the configuration for the rate limiter
* `test_client.go` and `test_rate_limiter.go` contain unit tests for the client and rate limiter respectively
* `Makefile` contains the build script for the project
* `.gitignore` contains the git ignore file for the project

### License

This project is licensed under the Apache License 2.0.

### Dependencies

The project depends on the following libraries:

* `github.com/coreos/etcd/etcdclient` for leader election and distributed locking
* `github.com/sirupsen/logrus` for logging

### Features

The project provides the following features:

* Leader election
* Distributed locking
* High-performance rate limiting
* Configurable quotas