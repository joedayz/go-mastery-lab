package main

import (
	"fmt"
	"runtime"
)

// ============================================================================
// BUILD TAGS
// ============================================================================
// Build tags permiten incluir/excluir código durante la compilación
// Ejecutar con: go build -tags dev
// ============================================================================

//go:build dev
// +build dev

func init() {
	fmt.Println("Running in DEV mode")
}

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}

