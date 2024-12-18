GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/publish cmd/publish/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/assign-hash cmd/assign-hash/main.go
