#!/bin/bash

# Get the current Git tag
GIT_TAG=$(git describe --tags --abbrev=0)

if [ -z "$GIT_TAG" ]; then
    echo "Error: No Git tag found."
    exit 1
fi

# Create the bin directory if it doesn't exist
mkdir -p bin

# Build for Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$GIT_TAG" -o bin/go-movies-$GIT_TAG-linux-amd64 main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-X main.version=$GIT_TAG" -o bin/go-movies-$GIT_TAG-windows-amd64.exe main.go

echo "Build complete. The executables are in the bin directory."
