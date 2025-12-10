package main

import (
	"strings"
	"testing"
)

// ============================================================================
// FUZZING EN GO (Go 1.18+)
// ============================================================================
// Fuzzing genera inputs aleatorios para encontrar bugs
// Ejecutar con: go test -fuzz=Fuzz
// ============================================================================

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func FuzzReverse(f *testing.F) {
	// Seed corpus (ejemplos iniciales)
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	// Fuzz target
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if strings.Contains(orig, "\xbd") {
			t.Skip() // Skip invalid UTF-8
		}
	})
}

