package main

import (
	"testing"
)

func Fib(n int, r bool) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1

	default:
		// Fibonacci done with recursion is very inefficient, generally recurssion is less efficient than for loop but has it's moments
		if r {
			return Fib(n-1, true) + Fib(n-2, true)
		}

		a, b := 0, 1

		for i := 1; i < n; i++ {
			a, b = b, a+b

		}

		return b
	}
}

func BenchmarkFib20T(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20, true)
	}
}

func BenchmarkFib20F(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20, false)
	}
}
