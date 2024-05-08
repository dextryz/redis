fmt:
	go mod tidy -compat=1.17
	gofmt -l -s -w .

server:
	go run ./cmd/server/main.go

client:
	go run ./cmd/client/main.go
