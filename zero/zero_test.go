// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package zero

import (
	"fmt"
	"testing"
)

func TestZero(t *testing.T) {
	var (
		_true         *bool      = True()
		_false        *bool      = False()
		_int          *int       = Int()
		_int32        *int32     = Int32()
		_int64        *int64     = Int64()
		_float32      *float32   = Float32()
		_float64      *float64   = Float64()
		_string       *string    = String()
		_struct       *struct{}  = Struct()
		_intArray     *[]int     = IntArray()
		_int32Array   *[]int32   = Int32Array()
		_int64Array   *[]int64   = Int64Array()
		_float32Array *[]float32 = Float32Array()
		_float64Array *[]float64 = Float64Array()
		_stringArray  *[]string  = StringArray()
	)
	fmt.Println(_true)
	fmt.Println(_false)
	fmt.Println(_int)
	fmt.Println(_int32)
	fmt.Println(_int64)
	fmt.Println(_float32)
	fmt.Println(_float64)
	fmt.Println(_string)
	fmt.Println(_struct)
	fmt.Println(_intArray)
	fmt.Println(_int32Array)
	fmt.Println(_int64Array)
	fmt.Println(_float32Array)
	fmt.Println(_float64Array)
	fmt.Println(_stringArray)

	fmt.Println()

	fmt.Println(*_true)
	fmt.Println(*_false)
	fmt.Println(*_int)
	fmt.Println(*_int32)
	fmt.Println(*_int64)
	fmt.Println(*_float32)
	fmt.Println(*_float64)
	fmt.Println(*_string)
	fmt.Println(*_struct)
	fmt.Println(*_intArray)
	fmt.Println(*_int32Array)
	fmt.Println(*_int64Array)
	fmt.Println(*_float32Array)
	fmt.Println(*_float64Array)
	fmt.Println(*_stringArray)
}
