// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sys

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/hollson/gdk/emoji"
)

const defaultErrMsg = "Error occur:"

var DefaultErrFormat = func(c *CallerInfo, format string, a ...interface{}) string {
	return fmt.Sprintf("%s %s:%d %s ", emoji.Error, c.Name, c.Line, fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...interface{}) error {
	c, err := Caller(2)
	if err != nil {
		return err
	}
	return errors.New(DefaultErrFormat(c, format, a))
}

/**
 * 用特定信息包装一个 error，使其包含代码堆栈信息
 * 如果 err 为空则返回空
 */
func WrapError(err error, wrapMsg string) error {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if !ok {
		return errors.New("WrapError 方法获取堆栈失败")
	}

	var wrapErr error = nil
	if err != nil {
		if wrapMsg == "" {
			wrapMsg = defaultErrMsg
		}
		errMsg := fmt.Sprintf("%s %s:%d Cause by: %s", wrapMsg, f.Name(), line, err.Error())
		wrapErr = errors.New(errMsg)
	}
	return wrapErr
}

/**
 * 用特定信息包装一个 error，使其包含代码堆栈信息
 * 如果 err 为空则返回空
 */
func WrapErrorf(err error, wrapMsgFmt string, a ...interface{}) error {
	wrapMsg := ""
	if wrapMsgFmt == "" {
		wrapMsg = defaultErrMsg
	} else {
		wrapMsg = fmt.Sprintf(wrapMsgFmt, a...)
	}

	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if !ok {
		return errors.New("WrapError 方法获取堆栈失败")
	}

	var wrapErr error = nil
	if err != nil {
		if wrapMsg == "" {
			wrapMsg = defaultErrMsg
		}
		errMsg := fmt.Sprintf("%s \n\tat %s:%d\nCause by: %s", wrapMsg, f.Name(), line, err.Error())
		wrapErr = errors.New(errMsg)
	}
	return wrapErr
}
