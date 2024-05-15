#!/bin/bash

# Create the bin directory if it doesn't exist
mkdir -p bin

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bin/main_linux main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/main_windows.exe main.go

echo "Build complete. The executables are in the bin directory."
