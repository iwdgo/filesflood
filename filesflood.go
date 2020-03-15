// Package filesflood attempts to reach the limits of the file system.
package filesflood

import (
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"syscall"
	"time"
)

// FilesFloodEmfile returns the number of errors triggered sorted by type: EMFILE, ENFILE, Temporary, Other
// Temporary is usually the sum of both previous types on Unix.
func FilesFloodEmfile(limit int) (emfile, enfile, temporary, others int) {
	var wg sync.WaitGroup
	for i := 0; i < 2*limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				f, err := ioutil.TempFile("", "emfileflood")
				if err != nil {
					var errno syscall.Errno
					if errors.As(err, &errno) && errno == 0 {
						// No error information. Check OS.
						err = getErrorFromOS()
						errors.As(err, &errno)
					}
					switch errno {
					case syscall.EMFILE:
						emfile++
					case syscall.ENFILE:
						enfile++
					default:
						if errno.Temporary() {
							temporary++
						} else {
							others++
						}
					}
					continue
				}
				defer os.Remove(f.Name())
				time.Sleep(time.Second) // Removal will wait for closing
				f.Close()
				return
			}
		}()
	}
	wg.Wait()
	return
}

// FilesFloodEnfile returns the number of errors triggered sorted by type: EMFILE, ENFILE, Temporary, Other
// Temporary is usually the sum of both previous types on Unix.
func FilesFloodEnfile() (emfile, enfile, temporary, others int) {

	for {
		f, err := ioutil.TempFile("", "enfileflood")
		if err != nil {
			var errno syscall.Errno
			if errors.As(err, &errno) && errno == 0 {
				err = getErrorFromOS()
				errors.As(err, &errno)
			}
			switch errno {
			case syscall.EMFILE:
				emfile++
			case syscall.ENFILE:
				// Exit on the first ENFILE error
				enfile++
				return
			default:
				if errno.Temporary() {
					temporary++
				}
				others++
			}
			continue
		}
		defer os.Remove(f.Name())
		time.Sleep(time.Microsecond)
		f.Close()
	}
}
