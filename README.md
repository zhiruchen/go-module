# go-module

```sh
~/Desktop> mkdir go-module
~/Desktop> cd go-module/
~/D/go-module> go mod init github.com/zhiruchen/go-module
~/D/go-module> go build ./...
go: finding go.uber.org/zap v1.9.1
go: downloading go.uber.org/zap v1.9.1
go: finding go.uber.org/atomic v1.3.2
go: finding go.uber.org/multierr v1.1.0
go: downloading go.uber.org/atomic v1.3.2
go: downloading go.uber.org/multierr v1.1.0
~/D/go-module> go run main.go
go module test
{"level":"info","ts":1535637509.09551,"caller":"common/log.go:12","msg":"failed to fetch URL","url":"https://google.com","attempt":3,"backoff":1}
```

