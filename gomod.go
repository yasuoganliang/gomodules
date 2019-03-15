package gomodules

import (
	"fmt"
)

// MyGoMod 导出方法
func MyGoMod(str string, version string) {
	fmt.Printf("this is %s version: %s\n", str, version)
}
