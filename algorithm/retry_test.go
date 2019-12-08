package algorithm

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	Retry(5, time.Second, foo)
	fmt.Println("done")
}

var n = 0

func foo() error {
	n++
	if n > 3 {
		return NoRetryError(errors.New("中断重试"))
	}
	fmt.Println(time.Now().Second())
	return errors.New("timeout error")
}

func TestNoRetryError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want stop
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoRetryError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoRetryError() = %v, want %v", got, tt.want)
			}
		})
	}
}
