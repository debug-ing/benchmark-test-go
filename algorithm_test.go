package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 7)
	}
}
