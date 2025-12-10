#!/bin/bash

# Script para compilar para múltiples plataformas

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64 main.go

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64.exe main.go

echo "Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 main.go

echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o bin/app-darwin-arm64 main.go

echo "Done! Binaries are in bin/"

