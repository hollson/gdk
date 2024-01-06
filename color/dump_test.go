//go:build !windows
// +build !windows

package color_test

import (
	"fmt"
	"testing"

	"github.com/hollson/gdk/color"
)

func ExampleText() {
	fmt.Println(color.Text("仅样式", color.Bold))
	fmt.Println(color.Text("仅字体", color.FontRed))
	fmt.Println(color.Text("仅背景", color.BackCyan))
	fmt.Println(color.Text("字体+背景", color.FontRed, color.BackYellow))
	fmt.Println(color.Text("样式+字体", color.Bold, color.FontBlue))
	fmt.Println(color.Text("样式+背景", color.Bold, color.BackYellow))
	fmt.Println(color.Text("样式+字体+背景", color.Bold, color.FontBlue, color.BackMagenta))
}

// 测试样式
func TestRender(t *testing.T) {
	ExampleText()
}

// func TestAbnormal(t *testing.T) {
// 	fmt.Println(color.Text("异样测试1", -1))
// 	fmt.Println(color.Text("异样测试2", 0))
// 	fmt.Println(color.Text("异样测试3", 10))
// 	fmt.Println(color.Text("异样测试4", 100))
// 	fmt.Println(color.Text("异样测试5", 100, 101))
// 	fmt.Println(color.Text("异样测试6", 100, 102, 103))
// 	fmt.Println(color.Text("异样测试7", 999))
// }

func TestTips(t *testing.T) {
	fmt.Println(color.Info("info test"))
	fmt.Println(color.Warn("warn test"))
	fmt.Println(color.Error("error test"))
	fmt.Println(color.Success("success test"))
	fmt.Println(color.Failure("failure test"))
}
