// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !RUNONLY
// +build !RUNONLY

package rat

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	// fmt.Printf("%v ", Fibonacci(-3))
	fmt.Printf("%v ", Fibonacci(0))
	fmt.Printf("%v ", Fibonacci(1))
	fmt.Printf("%v ", Fibonacci(2))
	fmt.Printf("%v ", Fibonacci(3))
	fmt.Printf("%v ", Fibonacci(4))
	fmt.Printf("%v ", Fibonacci(5))
	fmt.Printf("%v ", Fibonacci(6))
	fmt.Println()
}

func TestFibonacci46(t *testing.T) {
	Fibonacci(46)
}

//
// func TestFibonacci47(t *testing.T) {
// 	Fibonacci(47)
// }
//
// func TestFibonacci48(t *testing.T) {
// 	Fibonacci(48)
// }
//
// func TestFibonacci49(t *testing.T) {
// 	Fibonacci(49)
// }

// go test -bench=BenchmarkFibonacci -count=2
// go test -bench=^BenchmarkFibonacci$ -benchtime=2s -count=5
// go test -v fib_test.go -bench="^BenchmarkFibonacci$"  -run=none  -benchtime="2s" -count=3
// go test fib_test.go -bench="BenchmarkFibonacci"  -run=none  -benchtime="2s" -count=3
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40)
	}
}
