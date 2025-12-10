package main

import (
	"fmt"
	"testing"
)

// ============================================================================
// BENCHMARKS EN GO
// ============================================================================
// Los benchmarks miden el rendimiento del código
// Ejecutar con: go test -bench=.
// ============================================================================

func slowFunction(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func BenchmarkSlowFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowFunction(1000)
	}
}

// ============================================================================
// BENCHMARK CON DIFERENTES TAMAÑOS
// ============================================================================

func BenchmarkSlowFunctionSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slowFunction(size)
			}
		})
	}
}

// ============================================================================
// COMPARAR IMPLEMENTACIONES
// ============================================================================

func appendSlice(n int) []int {
	s := []int{}
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func preallocSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendSlice(1000)
	}
}

func BenchmarkPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		preallocSlice(1000)
	}
}

