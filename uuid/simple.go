// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid


// 缺省推荐，基于v4(基于随机数)
func New() UUID {
    v4, _ := NewV4()
    return v4
}
