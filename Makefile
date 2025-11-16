build:
	@go build -o bin/photon cmd/main.go

run: build
	@./bin/photon

test:
	@go test -v ./...
