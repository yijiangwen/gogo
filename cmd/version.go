package cmd

import (
	"fmt"
)

var version = "0.0.4"

// Version 打印版本号
func Version(args []string) {
	fmt.Printf("gogo version %s\n", version)
}

// VersionHelp 命令帮助
func VersionHelp(args []string) {
	fmt.Println(`
Usage: gogo version

print gogo version information.
	`)
}
