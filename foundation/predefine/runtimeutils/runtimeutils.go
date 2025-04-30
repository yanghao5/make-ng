package runtimeutils

import (
	"fmt"
	"runtime"
)

// i386, amd64, arm32, arm64, riscv64
func getRunTimeARCHBySTD() string {
	switch runtime.GOARCH {
	case "amd64":
		return "amd64"
	case "386":
		return "i386"
	case "arm":
		return "arm32"
	case "arm64":
		return "arm64"
	case "riscv64":
		return "riscv64"
	default:
		return ""
	}
}

func getOSPlatformBySTD() int {
	switch runtime.GOARCH {
	case "amd64":
		return 64
	case "386":
		return 32
	case "arm":
		return 32
	case "arm64":
		return 64
	case "riscv64":
		return 64
	default:
		return 0
	}
}

// `getOSTypeBySTD` get OS Type.
func getOSTypeBySTD() string {
	var platform int = getOSPlatformBySTD()
	switch runtime.GOOS {
	case "linux":
		return fmt.Sprintf("%s%d", runtime.GOOS, platform)
	case "windows":
		return fmt.Sprintf("%s%d", runtime.GOOS, platform)
	case "darwin":
		return "macOS"
	default:
		return ""
	}
}

func init() {
	ARCH = getRunTimeARCHBySTD()
	OS = getOSTypeBySTD()
}

var ARCH string

// Currently support windows64 windows32 linux32 linux64 macOS
// Subsequent support freebsd64 openbsd64 netbsd64
var OS string
