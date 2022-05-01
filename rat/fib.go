// Package rat:实验鼠
package rat

// Fibonacci 求斐波那契数列,模拟一个性能(耗时)运算
//  Test Outs:
//  Fibonacci(40) => 0.38s
//  Fibonacci(41) => 0.61s
//  Fibonacci(42) => 0.98s
//  Fibonacci(43) => 1.57s
//  Fibonacci(44) => 2.52s
//  Fibonacci(45) => 4.12s
//  Fibonacci(46) => 6.88s
//  Fibonacci(47) => 10.88s
//  Fibonacci(48) => 17.49s
//  Fibonacci(49) => 28.14s
//  Fibonacci(50) => 45.11s
//  Fibonacci(51) => 73.62s
//  Fibonacci(52) => 118.23s
func Fibonacci(num int) int {
	if num < 1 {
		return 0
	}
	if num < 3 {
		return 1
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
