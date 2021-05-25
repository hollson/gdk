// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rats


// 斐波那契数列(模拟一个耗时运算)
//  Test Outs:
//  fibonacci(40) => 0.38s
//  fibonacci(41) => 0.61s
//  fibonacci(42) => 0.98s
//  fibonacci(43) => 1.57s
//  fibonacci(44) => 2.52s
//  fibonacci(45) => 4.12s
//  fibonacci(46) => 6.88s
//  fibonacci(47) => 10.88s
//  fibonacci(48) => 17.49s
//  fibonacci(49) => 28.14s
//  fibonacci(50) => 45.11s
//  fibonacci(51) => 73.62s
//  fibonacci(52) => 118.23s
func Fibonacci(num int) int {
    if num < 1 {
        return 0
    }
    if num < 3 {
        return 1
    }
    return Fibonacci(num-1) + Fibonacci(num-2)
}