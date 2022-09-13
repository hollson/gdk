// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package inspect

import (
	"fmt"
	"os"
	"runtime"
	"strings"
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

// https://www.cnblogs.com/ucas123/p/14186171.html
type inspect struct {
	Name        string `json:"name"`        // App名称		"Awesome"
	Version     string `json:"version"`     // 应用版本		"v1.0.0"
	OSArch      string `json:"os_arch"`     // 系统架构:		runtime.GOARCH
	GoVersion   string `json:"go_version"`  // Go版本:		runtime.Version()
	Goos        string `json:"goos"`        // Go版本:		runtime.GOOS
	Branch      string `json:"branch"`      // Git分支:		git rev-parse --abbrev-ref @{u}
	Commit      string `json:"commit"`      // Git提交:		git rev-parse [--short] HEAD
	Author      string `json:"author"`      // Git用户名:	git config user.name
	Tag         string `json:"tag"`         // Git最近tag:	git describe --tags --abbrev=0
	Machine     string `json:"machine"`     // 主机名称:		bash hostname
	Environment string `json:"environment"` // 运行环境：		upper(${Name})_VERSION"
	Built       string `json:"built"`       // 构建时间		date "+%Y-%m-%d %H:%M:%S"
}

func (i inspect) String() string {
	format := `
	App Name:	 %s
	Version:	 %s
	OS/Arch:	 %s
	Go version:	 %s
	GOOS:	 %s
	Git Branch:	 %s
	Git Commit:	 %s
	Git Author:	 %s
	Git Tag:	 %s
	Machine:	 %s
	Environment:	 %s
	Built:	 %s
`
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

// Inspect 编译信息
func Inspect() inspect {
	ret := inspect{
		Name:        name,
		Version:     version,
		OSArch:      runtime.GOARCH,
		GoVersion:   runtime.Version(),
		Goos:        runtime.GOOS,
		Branch:      branch,
		Commit:      commit,
		Author:      author,
		Tag:         tag,
		Environment: environment,
		Built:       built,
	}
	ret.Machine, _ = os.Hostname()
	if len(ret.Environment) == 0 {
		ret.Environment = os.Getenv(fmt.Sprintf("%s_VERSION", strings.ToUpper(name)))
	}
	return ret
}

//go:generate go build -ldflags "-X main.buildVer=1.0"
//go:generate go build -ldflags "-X main.commitID=$(git rev-parse HEAD)" -o ./deploy/hello-server
//go:generate GOOS=linux GOARCH=amd64 go build -o ./deploy/hello-server
