package runtimeutils

import "runtime"

// i386, amd64, arm32, arm64, riscv64
func getRunTimeARCHbySTD() string {
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
		return runtime.GOARCH
	}
}

func init() {
	ARCH = getRunTimeARCHbySTD()
}

var ARCH string
