run:
	go run cmd/main.go cmd/server.go

build:
	GOOS=linux GOARCH=amd64 go build -o gateway cmd/main.go cmd/server.go
