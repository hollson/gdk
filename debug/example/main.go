package main

import (
	"fmt"
	"runtime"
)

/*
参考Docker：
$ docker version
Client:
  Cloud integration: v1.0.22
  Version:           20.10.12
  API version:       1.41
  Go version:        go1.16.12
  Git commit:        e91ed57
  Built:             Mon Dec 13 11:46:56 2021
  OS/Arch:           darwin/amd64
  Context:           default
  Experimental:      true
*/

var (
	version   string
	gitCommit string // Git提交: git rev-parse --short HEAD
	built     string // 编译时间: date "+%Y-%m-%d %H:%M:%S"
)

/*
运行:
go run  -ldflags \
   "-X 'main.version=v1.0.0'
   -X 'main.built=$(date "+%Y-%m-%d %H:%M:%S")'
   -X 'main.gitCommit=$(git rev-parse --short HEAD)'" \
  main.go
*/
func main() {
	fmt.Printf("\n版本信息: \n")
	fmt.Printf("Version: \t%s\n" +
		"Git commit: \t%s\n" +
		"OS/Arch: \t%s\n" +
		"Built: \t%s\n",
		version,gitCommit,runtime.GOARCH,built)
}
