// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"os"
)

var (
	_name    = "AWESOME"
	_version = "unknown"
	_branch  = "unknown"
	_commit  = "unknown"
	_author  = "unknown"
	_tag     = "unknown"
	_build   = "unknown"
)

// https://www.cnblogs.com/ucas123/p/14186171.html
type inspect struct {
	Name      string `json:"name"`       // App名称
	Branch    string `json:"branch"`     // git rev-parse --abbrev-ref @{u}
	Commit    string `json:"commit"`     // 最近的提交(git rev-parse [--short] HEAD)
	Author    string `json:"author"`     // Git用户名(git config user.name)
	Tag       string `json:"tag"`        // 最近的tag(git describe --tags --abbrev=0)
	Version   string `json:"version"`    // 自定义版本标识
	BuildTime string `json:"build_time"` // 构建时间
}

// Inspect 用于核查编译信息
//  go build -ldflags "-X debug.buildVer=1.0"
func Inspect() inspect {
	ret := inspect{
		Name:      _name,
		Branch:    _branch,
		Commit:    _commit,
		Author:    _author,
		Tag:       _tag,
		Version:   _version,
		BuildTime: _build,
	}
	if _version == "known" {
		ret.Version = os.Getenv(fmt.Sprintf("%s_VERSION", _name))
	}
	return ret
}

//go:generate go build -ldflags "-X main.buildVer=1.0"
//go:generate go build -ldflags "-X main.commitID=$(git rev-parse HEAD)" -o ./deploy/hello-server
//go:generate GOOS=linux GOARCH=amd64 go build -o ./deploy/hello-server
