// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"errors"
)

var (
	NotFound      = errors.New("not found")             // 未找到资源
	ParamError    = errors.New("param error")           // 参数错误
	AuthError     = errors.New("auth error")            // 鉴权失败
	IllegalUser   = errors.New("illegal user")          // 非法用户(跨账号操作等)
	CSRFFail      = errors.New("CSRF invalid")          // 跨站请求失败
	Forbidden     = errors.New("forbidden")             // 无权操作
	NetworkError  = errors.New("network error")         // 网络错误
	RepeatRequest = errors.New("repeat request")        // 重复请求(非幂等)
	OutdatedVer   = errors.New("outdated version")      // 版本过期/不兼容
	FileDataError = errors.New("file data error")       // 文件格式/编码错误
	FileSizeError = errors.New("file size error")       // 文件尺寸错误
	UploadFailed  = errors.New("upload failed")         // 文件上传失败
	GatewayError  = errors.New("gateway error")         // 网关错误
	ServerError   = errors.New("server error")          // 服务错误
	ServerTimeout = errors.New("server timeout")        // 服务处理超时
	InMaintenance = errors.New("in maintenance")        // 系统维护中
	DataOpError   = errors.New("data operation failed") // 数据(库)处理错误
)
