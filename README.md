# mockery example

```bash
GOBIN=$(git rev-parse --show-toplevel)/bin go install github.com/vektra/mockery/v2@v2.38.0
export PATH="$PWD/bin:$PATH"
go generate ./...
go test -v ./...
```

I also tried mockery v3

Check out other branches!
