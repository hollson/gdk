package async_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hollson/gdk/async"
)

func ExampleGo() {
	async.ExceptionHandler = func(err interface{}) {
		fmt.Printf(" ❌ goroutine error: %v\n", err)
	}

	async.Go(func() {
		panic("some panic")
	})

	time.Sleep(time.Second)

	// Output:
	// ❌ goroutine error: some panic
}

func TestGo(t *testing.T) {
	ExampleGo()
}
