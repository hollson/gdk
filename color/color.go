// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/*
 仅支持Linux，更过可参考https://github.com/gookit/color
*/

// Package color Linux控制台输出字体渲染
package color

import (
	"fmt"
)

// StyleOrColor 样式/颜色
type StyleOrColor uint8

const (
	Reset        StyleOrColor = 0 // 重置
	Bold                      = 1 // 加粗
	Faint                     = 2 // 模糊
	Italic                    = 3 // 斜体
	Underline                 = 4 // 下划线
	BlinkSlow                 = 5 // 慢速闪烁
	BlinkRapid                = 6 // 快速闪烁
	ReverseVideo              = 7 // 反白/反向显示
	Concealed                 = 8 // 隐藏/暗格
	CrossedOut                = 9 // 删除

	FontBlack   = 30 // 「字体」黑色
	FontRed     = 31 // 「字体」红色
	FontGreen   = 32 // 「字体」绿色
	FontYellow  = 33 // 「字体」黄色
	FontBlue    = 34 // 「字体」蓝色
	FontMagenta = 35 // 「字体」品红/洋紫
	FontCyan    = 36 // 「字体」青色
	FontWhite   = 37 // 「字体」白色

	BackBlack   = 40 // 「背景」黑色
	BackRed     = 41 // 「背景」红色
	BackGreen   = 42 // 「背景」绿色
	BackYellow  = 43 // 「背景」黄色
	BackBlue    = 44 // 「背景」蓝色
	BackMagenta = 45 // 「背景」品红/洋紫
	BackCyan    = 46 // 「背景」青色
	BackWhite   = 47 // 「背景」白色
)

// Text 渲染字体 (不同终端样式和颜色会有差异)
//  格式：color.Text("内容",样式,字体色,背景色)
//go:generate echo -e "\033[32m仅字体\033[0m"
//go:generate echo -e "\033[43m仅背景\033[0m"
//go:generate echo -e "\033[4m仅样式\033[0m"
//go:generate echo -e "\033[31;42m字体+背景\033[0m"
//go:generate echo -e "\033[1;33m字体+样式\033[0m"
//go:generate echo -e "\033[4;42m背景+样式\033[0m"
//go:generate echo -e "\033[4;37;45m字体+背景+样式\033[0m"
func Text(txt string, params ...StyleOrColor) string {
	_len := len(params)
	switch {
	case _len == 1:
		return fmt.Sprintf("\033[%dm%s\033[0m", params[0], txt)
	case _len == 2:
		return fmt.Sprintf("\033[%d;%dm%s\033[0m", params[0], params[1], txt)
	case _len > 2:
		return fmt.Sprintf("\033[%d;%d;%dm%s\033[0m", params[0], params[1], params[2], txt)
	default:
		return txt
	}
}
