#!/bin/bash

# Create dist directory if it doesn't exist
mkdir -p dist

# Build for Linux ARM
echo "Building for Linux ARM..."
export GOOS=linux
export GOARCH=arm
go build -o dist/api-linux-arm main.go
echo "Build completed for Linux ARM"

# Build for Linux AMD64
echo "Building for Linux AMD64..."
export GOOS=linux
export GOARCH=amd64
go build -o dist/api-linux-amd64 main.go
echo "Build completed for Linux AMD64"
