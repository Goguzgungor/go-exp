.PHONY: run build test fmt vet tidy clean

# Belirli bir komutu çalıştır: make run CMD=hello
CMD ?= hello

run:
	go run ./cmd/$(CMD)

build:
	mkdir -p bin
	go build -o bin/$(CMD) ./cmd/$(CMD)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf bin
