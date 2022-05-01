// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package deep

import (
	"reflect"
	"testing"
)

// reflect.DeepEqual用于比较array、slice、map、struct等深度一致
func TestDeepEqual(t *testing.T) {
	m1 := map[int]interface{}{1: []int{1, 2, 3}, 2: 3, 3: "a"}
	m2 := map[int]interface{}{1: []int{1, 2, 3}, 2: 3, 3: "a"}
	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("sorry")
	}
}
