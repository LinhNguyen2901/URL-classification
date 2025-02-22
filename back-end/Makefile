.PHONY: build clean

# Binary name
BINARY_NAME=bootstrap

# Build the project
build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) ./cmd/main.go
	zip function.zip $(BINARY_NAME)

# Clean build files
clean:
	rm -f $(BINARY_NAME)
	rm -f function.zip