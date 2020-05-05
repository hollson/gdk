//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-11-28
// @ Version: 1.0.0
//
// UUID: (Universally Unique Identifier) 通用唯一标识符
//-------------------------------------------------------------------------------------

package utility

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"io"
	"strings"
)

// 版本1：基于时间戳和MAC地址
func UUID() string {
	return uuid.NewV1().String()
}

// 版本2：基于POSIX UID/GID的DCE安全UUID
func UUID2(domain byte) string {
	return uuid.NewV2(domain).String()
}

// 版本3： 基于命名空间UUID和name的SHA-1散列
func UUID3(ns uuid.UUID, name string) string {
	return uuid.NewV3(ns, name).String()
}

// 版本4：基于随机数
func UUID4() string {
	return uuid.NewV4().String()
}

//版本5：基于命名空间UUID和name的SHA-1散列
func NewV5(ns uuid.UUID, name string) string {
	return uuid.NewV5(ns, name).String()
}

//生成Guid字串
func Guid() string {
	md5String := func(s string) string {
		h := md5.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	}

	box := make([]byte, 48)
	_, _ = io.ReadFull(rand.Reader, box)

	md5str := md5String(base64.URLEncoding.EncodeToString(box))
	return strings.ToUpper(md5str)
}
