package cmd

import "strings"

func Vendor(args []string) {

	_, exec := getPackageInfoAndExec(true)

	// `go get ${argv.join(' ')}`

	exec("go", "get", strings.Join(args, " "))

}
