package main

import (
	"errors"
	"fmt"
	"runtime"
)

// Redirect 终端重定向
type Redirect int

const (
	Stdin    Redirect = iota // 0:标准输入
	Stdout                   // 1:标准输出
	Stderr                   // 2:标准错误
	Pipeline                 // 3:管道输出(文件流等)
)

const defaultErrMsg = "Error occur:"

// Stdout
// pipeline

var DumpErrorFile = false // 是否打印go文件路径

func Errorf(format string, a ...interface{}) error {
	pc, fi, line, ok := runtime.Caller(1)
	if !ok {
		return errors.New("failed to get stack")
	}
	if !DumpErrorFile {
		fi = ""
	}
	fn := runtime.FuncForPC(pc)
	err := fmt.Sprintf("❌ %v %s:%d %s ", fi, fn.Name(), line, fmt.Sprintf(format, a...))
	return errors.New(err)
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

func main() {
	// DumpErrorFile = true
	fmt.Println(Errorf("发生错误:%s", "测试"))
	fmt.Println(WrapError(errors.New("AAA"), "BBBB"))
}
