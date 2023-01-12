package main

import (
	"fmt"

	"github.com/hollson/gdk/inspect"
)

var (
	name    string = "test"
	version string = "v1.0"
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
  -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'
  -X 'github.com/hollson/gdk/inspect.commit=$(git rev-parse --short HEAD)'
  -X 'github.com/hollson/gdk/inspect.author=$(git config user.name)'
  -X 'github.com/hollson/gdk/inspect.tag=$(git describe --tags --abbrev=0)'
  -X 'github.com/hollson/gdk/inspect.environment=develop'" \
  main.go
*/
func main() {
	// 项目私有元信息
	fmt.Printf("%s_%s\n\n", name, version)
	fmt.Println()

	// 通用编译信息
	fmt.Println("文本格式：")
	fmt.Println(inspect.Info())
	fmt.Println()

	fmt.Println("Json格式：")
	fmt.Println(inspect.Info().Json())
}
