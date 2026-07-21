go.mod
```markdown
module distributed-rate-limiter

go 1.19

require (
	github.com/coreos/etcd v0.0.0-20230223211515-0c3e2d1f8d9f
	github.com/sirupsen/logrus v1.9.0
)

replace (
	github.com/coreos/etcd v0.0.0-20230223211515-0c3e2d1f8d9f => github.com/coreos/etcd v3.4.20
)

```