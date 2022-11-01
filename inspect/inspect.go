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
	name        string
	version     string
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
type inspector struct {
	Name        string `json:"name,omitempty"`        // App名称 	 "Awesome"
	Version     string `json:"version,omitempty"`     // 应用版本 	 "v1.0.0"
	OSArch      string `json:"os_arch"`               // 系统架构 	 runtime.GOARCH
	GoVersion   string `json:"go_version"`            // Go版本 	 runtime.Version()
	Goos        string `json:"goos"`                  // Go版本 	 runtime.GOOS
	Branch      string `json:"branch,omitempty"`      // Git分支 	 git rev-parse --abbrev-ref @{u}
	Commit      string `json:"commit,omitempty"`      // Git提交 	 git rev-parse [--short] HEAD
	Author      string `json:"author,omitempty"`      // Git用户名  git config user.name
	Tag         string `json:"tag,omitempty"`         // Git最近tag  git describe --tags --abbrev=0
	Machine     string `json:"machine"`               // 主机名称 	 bash hostname
	Environment string `json:"environment,omitempty"` // 运行环境 	 upper(${Name})_VERSION"
	Built       string `json:"built,omitempty"`       // 构建时间 	 date "+%Y-%m-%d %H:%M:%S"
}

func (i *inspector) String() string {
	format := `App Name: %s
Version: %s
OS/Arch: %s
Go version: %s
GOOS: %s
Git Branch: %s
Git Commit: %s
Git Author: %s
Git Tag: %s
Machine: %s
Environment: %s
Built: %s`

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

func (i *inspector) Json() string {
	return gdk.JsonPretty(i)
}

/*
Info 检查应用编译信息
 go run -ldflags \
 "-X 'github.com/hollson/gdk/inspect.name=example'
  -X 'github.com/hollson/gdk/inspect.version=V1.0.0'
  -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'
  -X 'github.com/hollson/gdk/inspect.commit=$(git rev-parse --short HEAD)'
  -X 'github.com/hollson/gdk/inspect.author=$(git config user.name)'
  -X 'github.com/hollson/gdk/inspect.tag=$(git describe --tags --abbrev=0)'
  -X 'github.com/hollson/gdk/inspect.machine=$(hostname)'
  -X 'github.com/hollson/gdk/inspect.environment=develop'
  -X 'github.com/hollson/gdk/inspect.built=$(date "+%Y-%m-%d %H:%M:%S")'" \
  main.go

 参考Docker命令：
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
