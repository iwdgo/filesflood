package filesflood

import (
	"log"
	"runtime"
	"time"
)

// ExampleFilesFloodEmfile is triggering a files flood to test the maximum number of files
func ExampleFilesFloodEmfile() {
	// Unix's allow to set a very low limit but it seems ineffective on CI
	limit := 20
	switch runtime.GOOS {
	// Max number of threads on CI is 10000. It is reached before any limit on the file system
	case "windows":
		limit = 500
	}
	throttleFileSystem(limit)
	em, en, t, o := FilesFloodEmfile(limit)
	log.Printf("errors: EMFILE: %d, ENFILE: %d, temporary: %d, others: %d, ", em, en, t, o)
	// Output:
}

// ExampleFilesFloodEnfile is triggering a files flood to test the maximum number of files by thread
func ExampleFilesFloodEnfile() {
	timeout := 5 // in minutes
	// Unix's allow to set a very low limit but it seems ineffective on CI
	limit := 20
	switch runtime.GOOS {
	// Max number of threads on CI is 10000. It is reached before any limit on the file system
	case "windows":
		limit = 500
	}
	throttleFileSystem(limit)
	ch := make(chan int)
	go func() {
		em, en, t, o := FilesFloodEnfile()
		log.Printf("errors: EMFILE: %d, ENFILE: %d, temporary: %d, others: %d, ", em, en, t, o)
		ch <- 0
	}()

	select {
	case _ = <-ch:
	case <-time.After(time.Duration(timeout) * time.Minute):
		log.Printf("timed out of %d expired", timeout)
	}
	// Output:
}
