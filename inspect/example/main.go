package main

import (
	"fmt"

	"github.com/hollson/gdk/inspect"
)

var (
	name    string
	version string
)

/*
演示示例：
 构建:
 go build -ldflags \
 "-X 'main.name=example'
  -X 'main.version=V1.0.0'
  -X 'github.com/hollson/gdk/inspect.built=$(date "+%Y-%m-%d %H:%M:%S")'
  -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'" \
  -o ./example

 运行:
 go run -ldflags \
 "-X 'main.name=example'
  -X 'main.version=V1.0.0'
  -X 'github.com/hollson/gdk/inspect.built=$(date "+%Y-%m-%d %H:%M:%S")'
  -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'" \
  main.go
*/
func main() {
	fmt.Printf("%s_%s\n\n", name, version)
	fmt.Println(inspect.Info().Json())
}
