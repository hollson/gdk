package sys

import (
	"errors"
	"runtime"
	"strings"
)

// Organize 项目组织名称，如github.com/hollson/gdk中的hollson
var Organize = ""

type CallerInfo struct {
	PC   uintptr // 调用者函数指针
	Path string  // 调用者文件路径
	Name string  // 调用者函数名
	Line int     // 调用者行号
}

// Caller 获取调用者信息
func Caller(skips ...int) (*CallerInfo, error) {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}

	pc, fi, line, ok := runtime.Caller(skip)
	if !ok {
		return nil, errors.New("failed to get stack")
	}
	caller := CallerInfo{
		PC:   pc,
		Path: fi,
		Name: runtime.FuncForPC(pc).Name(),
		Line: line,
	}
	index := strings.Index(caller.Path, Organize)
	if len(Organize) > 0 && index > 0 {
		caller.Path = caller.Path[strings.Index(caller.Path, Organize):]
		caller.Name = caller.Name[strings.Index(caller.Name, Organize):]
	}
	return &caller, nil
}

// CallerMust 参考sys.Caller
func CallerMust() *CallerInfo {
	c, _ := Caller(2)
	return c
}
