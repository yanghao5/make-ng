package osutils

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func WinOS() (string, error) {
	type OSVersionInfoEx struct {
		OSVersionInfoSize uint32
		MajorVersion      uint32
		MinorVersion      uint32
		BuildNumber       uint32
		PlatformId        uint32
		CSDVersion        [128]uint16
	}

	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	getVersionEx := kernel32.NewProc("GetVersionExW")

	var osvi OSVersionInfoEx
	osvi.OSVersionInfoSize = uint32(unsafe.Sizeof(osvi))

	ret, _, err := getVersionEx.Call(uintptr(unsafe.Pointer(&osvi)))
	if ret == 0 {
		return "", fmt.Errorf("failed to get version: %v", err)
	}

	return fmt.Sprintf("Windows Version: %d.%d, Build: %d", osvi.MajorVersion, osvi.MinorVersion, osvi.BuildNumber), nil
}
