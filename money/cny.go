// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package money 货币单位/换算
package money

import (
	"fmt"
)

// CNY 人民币(单位：分)
type CNY int64

const (
	F CNY = 1         // 分
	J     = 10 * F    // 角
	Y     = 10 * J    // 元
	W     = 10000 * Y // 万
	E     = 10000 * W // 亿
)

// New 构造人民币对象
//  y: 元
//  j: 角
//  f: 分
func New(y, j, f CNY) CNY {
	return y*Y + j*J + f*F
}

func (c CNY) withUnit() string {
	switch {
	case c < J:
		return "分"
	case c < Y:
		return "角"
	case c < W:
		return "元"
	case c < E:
		return "万"
	default:
		return "亿"
	}
}

func (c CNY) dump() string {
	ret := ""
	switch {
	case c < J:
		ret = fmt.Sprintf("%.2f", float64(c)/float64(F))
	case c < Y:
		ret = fmt.Sprintf("%.2f", float64(c)/float64(J))
	case c < W:
		ret = fmt.Sprintf("%.2f", float64(c)/float64(Y))
	case c < E:
		ret = fmt.Sprintf("%.2f", float64(c)/float64(W))
	default:
		ret = fmt.Sprintf("%.2f", float64(c)/float64(E))
	}
	return ret
}

//Deprecated:
// Format 格式化输出
//  todo 如：time.Format("2006-01-02")
func (c CNY) Format() string {
	return c.String()
}

func (c CNY) String() string {
	return fmt.Sprintf("%s%s", c.dump(), c.withUnit())
}

// func (c CNY) Full() string {
// 	sb := strings.Builder{}
// 	if c >= E {
// 		_, _ = fmt.Fprintf(&sb, "%d亿", c/E)
// 		c %= E
// 	}
// 	if c >= W {
// 		_, _ = fmt.Fprintf(&sb, "%d万", c/W)
// 		c %= W
// 	}
// 	if c >= Y {
// 		_, _ = fmt.Fprintf(&sb, "%d元", c/Y)
// 		c %= Y
// 	}
//
// 	if c >= J {
// 		_, _ = fmt.Fprintf(&sb, "%d角", c/J)
// 		c %= J
// 	}
//
// 	if c >= F {
// 		_, _ = fmt.Fprintf(&sb, "%d分", c/F)
// 	}
// 	return sb.String()
// }
