package goGeek

import (
	"flag"
)

// 输入命令行参数
// 运行：go run go_02.go -name="rayest"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func go_02() {
	flag.Parse()
}
