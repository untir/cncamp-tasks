.PHONY: build

 build:
 	echo "building miniserver..."
 	mkdir -p bin/amd64
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 cmd/miniserver