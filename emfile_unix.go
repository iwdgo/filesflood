// +build !windows

package filesflood

import (
	"log"
	"syscall"
)

// throttleFileSystem lowers some file system limits
func throttleFileSystem(limit int) {
	// http://manpages.ubuntu.com/manpages/bionic/man2/getrlimit.2.html
	var rlimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		log.Fatal(err)
	}
	log.Printf("rlimits changed from: Cur = %d, Max = %d", rlimit.Cur, rlimit.Max)
	rlimit.Cur = uint64(limit / 2)
	rlimit.Max = uint64(limit)
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		log.Fatal(err)
	}
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		log.Fatal(err)
	}
	log.Printf("to: Cur = %d, Max = %d", rlimit.Cur, rlimit.Max)
}

func getErrorFromOS() error {
	return nil
}
