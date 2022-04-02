// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !go1.18
// +build !go1.18

package color

// 添加「编译约束」，参考： https://pkg.go.dev/cmd/go#hdr-Build_constraints
type any = interface{}
