// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package inspect

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/hollson/gdk"
)

var (
	name        string = "example"
	version     string = "v0.0.1"
	osArch      string
	goVersion   string
	goos        string
	branch      string
	commit      string
	author      string
	tag         string
	machine     string
	environment string
	built       string
)

// 加载运行时参数
func init() {
	osArch = runtime.GOARCH
	goVersion = runtime.Version()
	goos = runtime.GOOS
	machine, _ = os.Hostname()

	if len(environment) == 0 {
		environment = os.Getenv(fmt.Sprintf("%s_VERSION", strings.ToUpper(name)))
	}
}

// https://www.cnblogs.com/ucas123/p/14186171.html
// https://blog.51cto.com/cwind/654745
// https://blog.csdn.net/shelutai/article/details/127392916
// http://blog.champbay.com/2019/11/26/ldflags%E5%9C%A8golang%E7%BC%96%E8%AF%91%E4%B8%AD%E7%9A%842%E4%B8%AA%E4%BD%9C%E7%94%A8/
type inspector struct {
	Name        string `json:"name"`        // App名称: "Awesome"
	Version     string `json:"version"`     // 应用版本: "v1.0.0"
	OSArch      string `json:"os_arch"`     // 系统架构: runtime.GOARCH
	GoVersion   string `json:"go_version"`  // Go版本: runtime.Version()
	Goos        string `json:"goos"`        // Go版本: runtime.GOOS
	Branch      string `json:"branch"`      // Git分支: git rev-parse --abbrev-ref @{u}
	Commit      string `json:"commit"`      // Git提交: git rev-parse [--short] HEAD
	Author      string `json:"author"`      // Git用户名: git config user.name
	Tag         string `json:"tag"`         // Git最近tag: git describe --tags --abbrev=0
	Machine     string `json:"machine"`     // 主机名称: bash hostname
	Environment string `json:"environment"` // 运行环境: upper(${Name})_VERSION"
	Built       string `json:"built"`       // 构建时间: date "+%Y-%m-%d %H:%M:%S"
}

func (i *inspector) String() string {
	format := `Build Info:
  App Name: 	%s
  Version: 	%s
  OS/Arch: 	%s
  Go version: 	%s
  GOOS: 	%s
  Git Branch: 	%s
  Git Commit: 	%s
  Git Author: 	%s
  Git Tag: 	%s
  Machine: 	%s
  Environment: 	%s
  Built: 	%s`

	return fmt.Sprintf(format,
		name,
		version,
		osArch,
		goVersion,
		goos,
		branch,
		commit,
		author,
		tag,
		machine,
		environment,
		built)
}

// Json Json格式化
func (i *inspector) Json() string {
	return gdk.JsonPretty(i)
}

/*
Info 编译信息，即通过连接器参数将编译信息写入程序，参考 docker version 命令
 调用示例: fmt.Println(inspect.Info())
 # Output:
 Build Info:
  App Name:     awesome
  Version:      v1.0.0
  OS/Arch:      amd64
  Go version:   go1.17.8
  GOOS:         darwin
  Git Branch:   origin/dev
  Git Commit:   fedc48b
  Git Author:   hollson
  Git Tag:      list
  Machine:      shs
  Environment:  develop
  Built:        2022-12-01 15:04:05

 运行示例(编译时传入连接器参数):
 $ go run -ldflags \
 "-X 'github.com/hollson/gdk/inspect.name=awesome'
  -X 'github.com/hollson/gdk/inspect.version=v1.0.0'
  -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'
  -X 'github.com/hollson/gdk/inspect.commit=$(git rev-parse --short HEAD)'
  -X 'github.com/hollson/gdk/inspect.author=$(git config user.name)'
  -X 'github.com/hollson/gdk/inspect.tag=$(git describe --tags --abbrev=0)'
  -X 'github.com/hollson/gdk/inspect.machine=$(hostname)'
  -X 'github.com/hollson/gdk/inspect.environment=develop'
  -X 'github.com/hollson/gdk/inspect.built=$(date "+%Y-%m-%d %H:%M:%S")'" \
  main.go

 另一个推荐方案,是将编译信息写入version文件。
*/
func Info() *inspector {
	return &inspector{
		Name:        name,
		Version:     version,
		OSArch:      osArch,
		GoVersion:   goVersion,
		Goos:        goos,
		Branch:      branch,
		Commit:      commit,
		Author:      author,
		Tag:         tag,
		Machine:     machine,
		Environment: environment,
		Built:       built,
	}
}
