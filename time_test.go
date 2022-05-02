// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gdk

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	fmt.Println(ParseDate("2022-12-12"))
	fmt.Println(ParseTime("2022-12-12 22:22:22"))
}

func TestTime1970(t *testing.T) {
	fmt.Println(Time1970())
}

func Test1970(t *testing.T) {
	fmt.Println(time.UnixMicro(0))
}

func TestRFC3339(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "int, return error",
			input:   1,
			wantErr: true,
		},
		{
			name:    "Empty string",
			input:   "",
			wantErr: false,
		},
		{
			name:    "GMT +7",
			input:   "2019-10-12T14:20:50.52+07:00",
			wantErr: false,
		},
		{
			name:    "UTC",
			input:   "2019-10-12T14:20:50.52Z",
			wantErr: false,
		},
		{
			name:    "Broken string, return error",
			input:   "2019-10-12 14:20:50",
			wantErr: true,
		},
		{
			name:    "time.Time",
			input:   time.Now(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// gotErr := RFC3339(tt.input)
			// assert.Equal(t, tt.wantErr, gotErr != nil)
		})
	}
}
