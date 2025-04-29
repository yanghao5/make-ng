package main

import (
	"fmt"
	"makeng/foundation/osutils"
)

func main() {
	version, err := osutils.WinOS()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(version)
}
