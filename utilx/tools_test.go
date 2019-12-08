package utilx

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestThisMonday(t *testing.T) {
	fmt.Println(RetroWeekN(time.Tuesday))
}

func TestContainInt(t *testing.T) {
	type args struct {
		one int
		all []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt(tt.args.one, tt.args.all...); got != tt.want {
				t.Errorf("ContainInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainStr(t *testing.T) {
	type args struct {
		one string
		all []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainStr(tt.args.one, tt.args.all...); got != tt.want {
				t.Errorf("ContainStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuplicate(t *testing.T) {
	type args struct {
		a interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRet []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := Duplicate(tt.args.a); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Duplicate() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestStrToInt64(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt64(tt.args.num); got != tt.want {
				t.Errorf("StrToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSqlIn(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSqlIn(tt.args.items...); got != tt.want {
				t.Errorf("ToSqlIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestT3(t *testing.T) {
	type args struct {
		b bool
		x interface{}
		y interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T3(tt.args.b, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("T3() = %v, want %v", got, tt.want)
			}
		})
	}
}
