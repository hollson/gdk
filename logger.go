// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"log"
)

var Logger *log.Logger

type Loggerx interface {
	Trace(requestID string, fields map[string]interface{}, message string)
	Debug(requestID string, fields map[string]interface{}, message string)
	Info(requestID string, fields map[string]interface{}, message string)

	Warn(requestID string, err error, fields map[string]interface{}, message string)
	Error(requestID string, err error, fields map[string]interface{}, message string)
	Fatal(requestID string, err error, fields map[string]interface{}, message string)
}

func Logf(s string, args ...interface{}) {
	if Logger == nil {
		return
	}
	Logger.Output(2, fmt.Sprintf(s, args...))
}
