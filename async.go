// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"log"
)

// https://blog.csdn.net/qq_43158436/article/details/114582344

// ExceptionHandler 自定义错误处理程序
var ExceptionHandler = func(err interface{}) {
	log.Panicln(err)
}

// Go 异步函数
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				ExceptionHandler(err)
			}
		}()
		f()
	}()
}
