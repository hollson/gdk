// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package color

import (
	"fmt"
)

// Color 颜色
type FontColor int

// 前景色
const (
	FontBlack   FontColor = iota + 30 // 30: 黑色
	FontRed                           // 31: 红色
	FontGreen                         // 32: 绿色
	FontYellow                        // 33: 黄色
	FontBlue                          // 34: 蓝色
	FontMagenta                       // 35: 品红/洋紫
	FontCyan                          // 36: 青色
	FontWhite                         // 37: 白色
)

// BackGroundColor

type BackGroundColor int

// 背景色
const (
	BackBlack   BackGroundColor = iota + 40 // 40: 黑色
	BackRed                                 // 41: 红色
	BackGreen                               // 42: 绿色
	BackYellow                              // 43: 黄色
	BackBlue                                // 44: 蓝色
	BackMagenta                             // 45: 品红/洋紫
	BackCyan                                // 46: 青色
	BackWhite                               // 47: 白色
)

// Style 样式
type Style int

const (
	Reset        Style = iota // 0: 重置
	Bold                      // 1: 加粗
	Faint                     // 2: 模糊
	Italic                    // 3: 斜体
	Underline                 // 4: 下划线
	BlinkSlow                 // 5: 慢速闪烁
	BlinkRapid                // 6: 快速闪烁
	ReverseVideo              // 7: 反白/反向显示
	Concealed                 // 8: 隐藏/暗格
	CrossedOut                // 9: 删除
)

// Text 渲染字体
//  格式：color.Text("内容",样式,字体色,背景色)
//go:generate echo -e "\033[32m仅字体\033[0m"
//go:generate echo -e "\033[43m仅背景\033[0m"
//go:generate echo -e "\033[4m仅样式\033[0m"
//go:generate echo -e "\033[31;42m字体+背景\033[0m"
//go:generate echo -e "\033[1;33m字体+样式\033[0m"
//go:generate echo -e "\033[4;42m背景+样式\033[0m"
//go:generate echo -e "\033[4;37;45m字体+背景+样式\033[0m"
func Text(txt string, params ...any) string {
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
