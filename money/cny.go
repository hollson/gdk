// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package money

import "fmt"

// 人民币(单位：分)
type CNY float64

const (
    F CNY = 1         // 分
    J     = 10 * F    // 角
    Y     = 10 * J    // 元
    Q     = 1000 * Y  // 千
    W     = 10 * Q    // 万
    E     = 10000 * W // 亿
)

// 转换为带两位小数的元
func (c CNY) Float64() float64 {
    return float64(c)
}

// 格式化输出(四舍五入),不带单位
func (c CNY) Single() string {
    ret := ""
    switch {
    case c < J:
        ret = fmt.Sprintf("%.2f", float64(c)/float64(F))
    case c < Y:
        ret = fmt.Sprintf("%.2f", float64(c)/float64(J))
    case c < Q:
        ret = fmt.Sprintf("%.2f", float64(c)/float64(Y))
    case c < W:
        ret = fmt.Sprintf("%.2f", float64(c)/float64(Q))
    case c < E:
        ret = fmt.Sprintf("%.2f", float64(c)/float64(W))
    default:
        ret = fmt.Sprintf("%.2f", float64(c)/float64(E))
    }
    return ret
}

// 格式化输出(四舍五入),带单位
func (c CNY) String() string {
    return fmt.Sprintf("%s%s", c.Single(), c.withUnit())
}

func (c CNY) withUnit() string {
    switch {
    case c < J:
        return "分"
    case c < Y:
        return "角"
    case c < Q:
        return "元"
    case c < W:
        return "千元"
    case c < E:
        return "万元"
    default:
        return "亿元"
    }
}
