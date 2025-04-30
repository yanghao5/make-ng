package osutils

import (
	"fmt"
	"log"
	"runtime"

	"golang.org/x/sys/windows"
)

// `WinOS` get windows version
func WinOS() string {
	if runtime.GOOS != "windows" {
		log.Printf("\033[31merror\033[0m: WinOS() unsupported OS: %s, only Windows is supported", runtime.GOOS)
		return ""
	}
	maj, min, patch := windows.RtlGetNtVersionNumbers()
	return fmt.Sprintf("%d.%d.%d", maj, min, patch)
}
