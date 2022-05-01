// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

type Request struct {
}

type Result struct {
	Code int
	Msg  string
	Data interface{}
}
