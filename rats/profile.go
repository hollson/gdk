// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rats


// 斐波那契数列(模拟一个耗时运算)
// Test Outs:
// fibonacci(42) => 981.471912ms
// fibonacci(43) => 1.569058476s
// fibonacci(49) => 28.107765764s
// fibonacci(50) => 45.106843729s
// fibonacci(51) => 73.621611718s
// fibonacci(52) => 118.23563034s
func Fibonacci(num int) int {
    if num < 1 {
        return 0
    }
    if num < 3 {
        return 1
    }
    return Fibonacci(num-1) + Fibonacci(num-2)
}