package main

import (
	"fmt"
	"makeng/foundation/osutils"
)

func main() {
	version, _ := osutils.WinOS()
	fmt.Println(version)
}
