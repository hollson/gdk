// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")             // 未找到资源
	ErrParam         = errors.New("param error")           // 参数错误
	ErrAuth          = errors.New("auth error")            // 鉴权失败
	ErrIllegalUser   = errors.New("illegal user")          // 非法用户(跨账号操作等)
	ErrCSRFFail      = errors.New("CSRF invalid")          // 跨站请求失败
	ErrForbidden     = errors.New("forbidden")             // 无权操作
	ErrNetwork       = errors.New("network error")         // 网络错误
	ErrRepeatRequest = errors.New("repeat request")        // 重复请求(非幂等)
	ErrOutdatedVer   = errors.New("outdated version")      // 版本过期/不兼容
	ErrFileData      = errors.New("file data error")       // 文件格式/编码错误
	ErrFileSize      = errors.New("file size error")       // 文件尺寸错误
	ErrUploadFailed  = errors.New("upload failed")         // 文件上传失败
	ErrGateway       = errors.New("gateway error")         // 网关错误
	ErrServer        = errors.New("server error")          // 服务错误
	ErrServerTimeout = errors.New("server timeout")        // 服务处理超时
	ErrInMaintenance = errors.New("in maintenance")        // 系统维护中
	ErrDataOperation = errors.New("data operation failed") // 数据(库)处理错误
)
