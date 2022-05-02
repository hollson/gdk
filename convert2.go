package gdk

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// ERR_NEED_NUMERIC
var ERR_NEED_NUMERIC = errors.New("ToInt64 need numeric")

// ERR_BYTES_INVALILD
var ERR_BYTES_INVALILD = errors.New("BytesToFloat64 bytes invalid")

// BytesToFloat64 convert bytes to float64
func BytesToFloat64(bytes []byte) (data float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ERR_BYTES_INVALILD
		}
	}()
	data = math.Float64frombits(binary.LittleEndian.Uint64(bytes))
	return data, err
}

// Float64ToBytes convert float64 to bytes; []uint8
func Float64ToBytes(input float64) []byte {
	bits := math.Float64bits(input)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// Float64ToStr convert float64 to string 支持指定精度
func Float64ToStr(num float64, precision int) string {
	return strconv.FormatFloat(num, 'f', precision, 64)
}

// StrToFloat64 convert string to float64, supported the given precision
func StrToFloat64(str string, precision int) float64 {
	precfmt := fmt.Sprintf("%%.%df", precision)
	value, _ := strconv.ParseFloat(str, 64)
	nstr := fmt.Sprintf(precfmt, value) // 指定精度
	val, _ := strconv.ParseFloat(nstr, 64)
	return val
}

// StrToFloat64Round  convert string to float64, supported the given precision and round
func StrToFloat64Round(str string, precision int, round bool) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return Float64Precision(f, precision, round)
}

// Float64Precision float指定精度; round为true时, 表示支持四舍五入
func Float64Precision(f float64, precision int, round bool) float64 {
	pow10N := math.Pow10(precision)
	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}
	return math.Trunc((f)*pow10N) / pow10N
}

// StructToMap struct convert to map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

// MapToStruct map obj to struct data
func MapToStruct(obj map[string]interface{}, data interface{}) (interface{}, error) {
	for k, v := range obj {
		err := setField(data, k, v)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structVal := reflect.ValueOf(obj).Elem()
	structFieldVal := structVal.FieldByName(name)

	if !structFieldVal.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldVal.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldVal.Type()
	if value == nil {
		return nil
	}
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match obj field type")
	}
	structFieldVal.Set(val)
	return nil
}
