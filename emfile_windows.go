// +build windows

package filesflood

import "syscall"

func throttleFileSystem(limit int) {
	// Not done programmatically. Check gowindows github/Action
}

func getErrorFromOS() error {
	return syscall.GetLastError()
}
