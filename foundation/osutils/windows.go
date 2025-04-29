package osutils

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

func WinOS() (string, error) {
	if runtime.GOOS != "windows" {
		return "", fmt.Errorf("\033[31merror\033[0m: WinOS() unsupported OS: %s, only Windows is supported", runtime.GOOS)
	}
	maj, min, patch := windows.RtlGetNtVersionNumbers()
	return fmt.Sprintf("%d.%d.%d", maj, min, patch), nil
}
