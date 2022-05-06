// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sys

// Redirect 终端重定向
type Redirect int

const (
	Stdin    Redirect = iota // 0:标准输入
	Stdout                   // 1:标准输出
	Stderr                   // 2:标准错误
	Pipeline                 // 3:管道输出(文件流等)
)
