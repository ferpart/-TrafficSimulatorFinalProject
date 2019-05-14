# Makefile taken from https://sohlich.github.io/post/go_makefile/
# Basic go commands

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME= CityTraffic
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	$(GOGET) -v
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GOGET) -v
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
