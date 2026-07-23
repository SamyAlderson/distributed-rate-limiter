# Distributed Rate Limiter
A high-performance, low-latency rate limiter for scalable systems.

## What it does
This is a simple, distributed rate limiter designed to handle high traffic volumes with minimal latency. It's intended for use in scalable systems where rate limiting is crucial.

## Install
To install the distributed rate limiter, run:
```bash
go get -u github.com/SamyAlderson/distributed-rate-limiter
```
## Usage
To use the distributed rate limiter, import it in your Go program and create a new instance:
```go
import (
	"github.com/SamyAlderson/distributed-rate-limiter"
)

func main() {
	rateLimiter := distributedrate.NewRateLimiter("my_key")
	// ...
}
```
## Build from source
To build the distributed rate limiter from source, run:
```bash
go build github.com/SamyAlderson/distributed-rate-limiter
```
## Run tests
To run the tests, navigate to the project directory and run:
```bash
go test github.com/SamyAlderson/distributed-rate-limiter
```
## Project structure
* `distributed_rate_limiter.go`: The main rate limiter implementation
* `test_distributed_rate_limiter.go`: The test suite for the rate limiter
* `README.md`: This file
* `LICENSE`: The MIT license

## License
Copyright (c) 2026 SamyAlderson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.