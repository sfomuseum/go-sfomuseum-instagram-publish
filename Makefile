cli:
	go build -mod vendor -o bin/publish cmd/publish/main.go
	go build -mod vendor -o bin/assign-hash cmd/assign-hash/main.go
