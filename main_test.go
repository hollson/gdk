// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"testing"

	"github.com/hollson/goox/value"
)

func TestDefaultValue(t *testing.T) {
	var a *string = &value.String
	var b *int64 = &value.Int64
	log.Println("default string pointer",a)
	log.Println("default int64 pointer",b)
}
