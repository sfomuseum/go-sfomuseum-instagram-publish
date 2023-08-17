GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/publish cmd/publish/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/assign-hash cmd/assign-hash/main.go
