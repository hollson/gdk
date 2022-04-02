// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package color

import (
	"fmt"
	"testing"
)

func ExampleText() {
	fmt.Println(Text("仅字体", FontRed))
	fmt.Println(Text("仅背景", BackCyan))
	fmt.Println(Text("仅样式", Bold)) // 粗体
	fmt.Println(Text("字体+背景", FontRed, BackYellow))
	fmt.Println(Text("字体+样式", Bold, FontBlue))
	fmt.Println(Text("背景+样式", Bold, BackYellow))
	fmt.Println(Text("字体+背景+样式", Bold, FontBlue, BackMagenta))
}

// 测试样式
func TestRender(t *testing.T) {
	ExampleText()
}

func TestAbnormal(t *testing.T) {
	fmt.Println(Text("异样测试1", -1))
	fmt.Println(Text("异样测试2", 0))
	fmt.Println(Text("异样测试3", 10))
	fmt.Println(Text("异样测试4", 100))
	fmt.Println(Text("异样测试5", 100, 101))
	fmt.Println(Text("异样测试6", 100, 102, 103))
	fmt.Println(Text("异样测试7", 999))
}
