#!/bin/bash

# Get the current Git commit hash
GIT_HASH=$(git rev-parse --short HEAD)

if [ -z "$GIT_HASH" ]; then
    echo "Error: No Git commit hash found."
    exit 1
fi

# Create the bin directory if it doesn't exist
mkdir -p bin

# Build for Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$GIT_HASH" -o bin/go-movies-$GIT_HASH-linux-amd64 main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-X main.version=$GIT_HASH" -o bin/go-movies-$GIT_HASH-windows-amd64.exe main.go

echo "Build complete. The executables are in the bin directory."