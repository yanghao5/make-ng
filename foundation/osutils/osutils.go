package osutils

import (
	"errors"
	"fmt"
)

// VersionInfo holds info about the linux kernel.
type VersionInfo struct {
	Major  int    // Major part of the kernel (e.g. 6.15.1-generic -> 6)
	Minor  int    // Minor part of the kernel version (e.g. 6.15.1-generic -> 15)
	Patch  int    // Patch part of the kernel version (e.g. 6.15.1-generic -> 1)
	Flavor string // Flavor of the kernel version (e.g. 6.15.1-generic -> generic)
}

func (k *VersionInfo) ToString() string {
	return fmt.Sprintf("%d.%d.%d%s", k.Major, k.Minor, k.Patch, k.Flavor)
}

// CompareKernelVersion compares two kernel.VersionInfo structs.
// Returns -1 if a < b
// 0 if a == b
// 1 if a > b
func CompareKernelVersion(a, b VersionInfo) int {
	if a.Major < b.Major {
		return -1
	} else if a.Major > b.Major {
		return 1
	}

	if a.Minor < b.Minor {
		return -1
	} else if a.Minor > b.Minor {
		return 1
	}

	if a.Patch < b.Patch {
		return -1
	} else if a.Patch > b.Patch {
		return 1
	}

	return 0
}

// ParseRelease parses a string and creates a VersionInfo based on it.
func ParseRelease(release string) (*VersionInfo, error) {
	var (
		major, minor, patch, parsed int
		flavor, partial             string
	)

	// Ignore error from Sscanf to allow an empty flavor.  Instead, just
	// make sure we got all the version numbers.
	parsed, _ = fmt.Sscanf(release, "%d.%d%s", &major, &minor, &partial)
	if parsed < 2 {
		return nil, errors.New("error: can't parse kernel version " + release)
	}

	// sometimes we have 3.12.25-gentoo, but sometimes we just have 3.12-1-amd64
	parsed, _ = fmt.Sscanf(partial, ".%d%s", &patch, &flavor)
	if parsed < 1 {
		flavor = partial
	}

	return &VersionInfo{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Flavor: flavor,
	}, nil
}
