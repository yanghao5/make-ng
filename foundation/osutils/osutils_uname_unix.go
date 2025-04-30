//go:build linux || freebsd || openbsd || netbsd

package osutils

import "golang.org/x/sys/unix"

func uname() (*unix.Utsname, error) {
	uts := &unix.Utsname{}

	if err := unix.Uname(uts); err != nil {
		return nil, err
	}
	return uts, nil
}
