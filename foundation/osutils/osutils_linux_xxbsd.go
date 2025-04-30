//go:build linux || freebsd || openbsd

package osutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

// https://github.com/moby/moby/blob/v28.1.1/pkg/parsers/kernel/kernel_unix.go
// GetKernelVersion gets the current kernel version.
// version,_:=osutils.GerKernelVersion()
// version.ToString()
func GetKernelVersion() (*VersionInfo, error) {
	uts, err := uname()
	if err != nil {
		return nil, err
	}

	// Remove the \x00 from the release for Atoi to parse correctly
	return ParseRelease(unix.ByteSliceToString(uts.Release[:]))
}

// CheckKernelVersion checks if current kernel is newer than (or equal to)
// the given version.
func CheckKernelVersion(major, minor, patch int) bool {
	if v, err := GetKernelVersion(); err != nil {
		log.Printf("error: getting kernel version: %s", err)
	} else {
		if CompareKernelVersion(*v, VersionInfo{Major: major, Minor: minor, Patch: patch}) < 0 {
			return false
		}
	}
	return true
}

func LinuxOS() string {
	v, err := GetKernelVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func BsdOS() string {
	v, err := GetKernelVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func GetLinuxDistro() (string, string) {
	// Open /etc/os-release file to read the OS information
	file, err := os.Open("/etc/os-release")
	if err != nil {
		log.Println("\033[31merror\033[0m: /etc/os-release ", err)
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var id, versionID string

	// Iterate through the file lines to extract ID and VERSION_ID
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			id = strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
		}
		if strings.HasPrefix(line, "VERSION_ID=") {
			versionID = strings.Trim(strings.TrimPrefix(line, "VERSION_ID="), "\"")
		}
	}

	// Handle any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading /etc/os-release:", err)
	}

	// Return the ID and VERSION_ID
	return id, versionID
}

// `LinuxDistro` returns the distribution name and version, based on /etc/os-release.
// The following distributions have been tested:
// - Debian         debian      8-13
// - Ubuntu         ubuntu      16.04-25.04
// - Fedora         fedora      20-43
// - CentOS         centos      7.9-8+
// - Rocky          rocky      <version-range>
// - Alma           alma       <version-range>
// - Alpine         alpine      3.1-3.21.3
// - openSUSE       opensuse    15.6
// - Arch Linux     arch        rolling release
// - VMware Photon OS photon     3.0-5.0
// - Fedora CoreOS  fedora-coreos <version-range>
// - Ubuntu Core    ubuntu-core <version-range>
// - Container-Optimized OS <os-name> <version-range>
// - Gentoo Linux   gentoo      2.17
// - Oracle Linux   ol          7.2-9.5
// - Kali           kali        <version-range>
// - Mageia         mageia      5-9
// - Linux Mint     linuxmint   6+
// - SLES           sles        12.5
// - Amazon Linux   amzn        2+
// - Clear Linux OS clear-linux-os 43300+
// - AltLinux       altlinux    p9-p10

// Untested distros might encounter issues when using the `LinuxDistro` function.
// Some distros don't have `/etc/os-release`, such as NixOS, RHEL, and Manjaro.
// So they may require additional handling.

func LinuxDistro() (string, string) {
	return GetLinuxDistro()
}
