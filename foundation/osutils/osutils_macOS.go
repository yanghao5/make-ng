//go:build darwin

package osutils

import (
	"fmt"
	"os/exec"
	"strings"
)

// `MacOS` get macOS system version
func MacOS() string {
	osName, err := getSPSoftwareDataType()
	if err != nil {
		return ""
	}
	v, err := getMacOSRelease(osName)
	if err != nil {
		return ""
	}
	return v
}

func Darwin() string {
	v, _ := getKernelVersion()
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// GetKernelVersion gets the current kernel version.
func getKernelVersion() (*VersionInfo, error) {
	osName, err := getSPSoftwareDataType()
	if err != nil {
		return nil, err
	}
	release, err := getDarwinRelease(osName)
	if err != nil {
		return nil, err
	}
	return ParseRelease(release)
}

// `getRelease` uses `system_profiler SPSoftwareDataType` to get macOS version
func getMacOSRelease(osName string) (string, error) {
	for _, line := range strings.Split(osName, "\n") {
		if !strings.Contains(line, "System Version") {
			continue
		}
		// It has the format like '      System Version: macOS 13.5.2 (22G91)'
		_, ver, ok := strings.Cut(line, "macOS")
		if !ok {
			return "", fmt.Errorf("error: parse macOS version")
		}

		ver = strings.Split(ver, "(")[0]
		var release string = strings.TrimSpace(ver)
		return release, nil
	}

	return "", nil
}

// `getRelease` uses `system_profiler SPSoftwareDataType` to get Darwin kernel version
func getDarwinRelease(osName string) (string, error) {
	for _, line := range strings.Split(osName, "\n") {
		if !strings.Contains(line, "Kernel Version") {
			continue
		}
		// It has the format like '      Kernel Version: Darwin 22.6.0'
		_, ver, ok := strings.Cut(line, ":")
		if !ok {
			return "", fmt.Errorf("kernel Version is invalid")
		}

		_, release, ok := strings.Cut(strings.TrimSpace(ver), " ")
		if !ok {
			return "", fmt.Errorf("kernel version needs to be 'Darwin x.x.x'")
		}
		return release, nil
	}

	return "", nil
}

func getSPSoftwareDataType() (string, error) {
	cmd := exec.Command("system_profiler", "SPSoftwareDataType")
	osName, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(osName), nil
}
