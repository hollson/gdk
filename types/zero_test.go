// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"testing"
)

func TestZero(t *testing.T) {
	var (
		_true         = True()
		_false        = False()
		_int          = Int()
		_int32        = Int32()
		_int64        = Int64()
		_float32      = Float32()
		_float64      = Float64()
		_string       = String()
		_struct       = Struct()
		_intArray     = IntArray()
		_int32Array   = Int32Array()
		_int64Array   = Int64Array()
		_float32Array = Float32Array()
		_float64Array = Float64Array()
		_stringArray  = StringArray()
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

func TestTrue(t *testing.T) {
	var a = True()
	var b = True()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestFalse(t *testing.T) {
	var a = False()
	var b = False()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestFloat32(t *testing.T) {
	var a = Float32()
	var b = Float32()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestFloat32Array(t *testing.T) {
	var a = Float32Array()
	var b = Float32Array()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}

	*a = append(*a, 1)
	fmt.Println(*a, *b)
}

func TestFloat64(t *testing.T) {
	var a = Float64()
	var b = Float64()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestFloat64Array(t *testing.T) {
	var a = Float64Array()
	var b = Float64Array()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}

	*a = append(*a, 1)
	fmt.Println(*a, *b)
}

func TestInt(t *testing.T) {
	var a = Int()
	var b = Int()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestInt32(t *testing.T) {
	var a = Int32()
	var b = Int32()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestInt32Array(t *testing.T) {
	var a = Int32Array()
	var b = Int32Array()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}

	*a = append(*a, 1)
	fmt.Println(*a, *b)
}

func TestInt64(t *testing.T) {
	var a = Int64()
	var b = Int64()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestInt64Array(t *testing.T) {
	var a = Int64Array()
	var b = Int64Array()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}

	*a = append(*a, 1)
	fmt.Println(*a, *b)
}

func TestIntArray(t *testing.T) {
	var a = IntArray()
	var b = IntArray()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}

	*a = append(*a, 1)
	fmt.Println(*a, *b)
}

func TestString(t *testing.T) {
	var a = String()
	var b = String()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}
}

func TestStringArray(t *testing.T) {
	var a = StringArray()
	var b = StringArray()

	if a == b {
		t.Error(a, b)
		t.Error(*a, *b)
	}

	*a = append(*a, "a")
	fmt.Println(*a, *b)
}

func TestStruct(t *testing.T) {
	var a = Struct()
	var b = Struct()
	b = &struct{}{}

	if a != b {
		t.Errorf("%p,%p", a, b)
	}
}
