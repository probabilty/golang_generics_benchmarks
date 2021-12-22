package generics

import (
	"math/rand"
	"testing"
)

func BenchmarkAddInt64Generic(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		addGeneric(rand.Int63(), rand.Int63())
	}
}

func BenchmarkAddInt64(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		addInt64(rand.Int63(), rand.Int63())
	}
}

func BenchmarkAddaddInterface(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		addInterface(rand.Int63(), rand.Int63())
	}
}
func BenchmarkFibonacciGeneric(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciGeneric(20)
	}
}

func BenchmarkFibonacciInt(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciInt64(20)
	}
}

func BenchmarkFibonacciInterface(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciInterface(20)
	}
}
