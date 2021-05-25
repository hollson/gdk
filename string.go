// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"regexp"
	"strings"
)

// 修剪所有换行和空格，如：压缩整段格式化的SQL
func TrimSpace(str string) string {
	if len(str) == 0 {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return strings.TrimSpace(reg.ReplaceAllString(str, " "))
}
