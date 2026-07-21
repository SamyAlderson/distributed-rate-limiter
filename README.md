# distributed-rate-limiter
A high-performance, distributed rate limiter for scalable systems

## What it does
This is a Go library for building scalable rate limiting systems. It uses leader election and distributed locking to ensure that rate limits are enforced consistently across multiple nodes. The library is highly configurable, allowing you to set up custom quotas and rate limits for your application.

## Installation
To use this library, run the following command:
```bash
go get github.com/SamyAlderson/distributed-rate-limiter
```
## Usage
To use the distributed rate limiter, import the library and create a new instance:
```go
import (
	"github.com/SamyAlderson/distributed-rate-limiter"
)

func main() {
	limiter := distributedrate.NewLimiter("my-namespace", 10, time.Minute)
	// use the limiter to rate limit requests
}
```
## Building from source
To build the library from source, run the following command:
```bash
go build github.com/SamyAlderson/distributed-rate-limiter
```
## Running tests
To run the tests, use the following command:
```bash
go test github.com/SamyAlderson/distributed-rate-limiter
```
## Project structure
* `cmd/`: contains the main executable
* `pkg/`: contains the rate limiter package
* `test/`: contains the test suite
* `main.go`: the main entry point for the library
* `limiter.go`: implements the rate limiter logic
* `leader_election.go`: implements leader election
* `distributed_lock.go`: implements distributed locking
* `config.go`: defines the config structure
* `tests.go`: contains the test suite

## License
Copyright (c) 2026 SamyAlderson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.