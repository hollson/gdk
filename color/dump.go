//go:build !windows
// +build !windows

package color

import (
	"fmt"

	"github.com/hollson/gdk/emoji"
)

// Text 精简快捷的终端彩色打印(仅支持linux)
//  格式：color.Text("内容",样式,字体色,背景色)
//  推荐使用: https://github.com/gookit/color
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

// Info 🔔
func Info(txt string) string {
	return Text(fmt.Sprintf(" %s %s", emoji.Info, txt), FontBlue)
}

// Warn ⚠️
func Warn(txt string) string {
	return Text(fmt.Sprintf(" ⚠️ %s", txt), FontMagenta)
}

// Error ❌
func Error(txt string) string {
	return Text(fmt.Sprintf(" %s %s", emoji.Error, txt), Bold, FontRed)
}

// Success ✅
func Success(txt string) string {
	return Text(fmt.Sprintf(" %s %s", emoji.Ok, txt), Bold, FontGreen)
}

// Failure 🚫
func Failure(txt string) string {
	return Text(fmt.Sprintf(" %s %s", emoji.No, txt), Bold, FontRed)
}
