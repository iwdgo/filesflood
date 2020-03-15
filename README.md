[![Go Reference](https://pkg.go.dev/badge/iwdgo/filesflood.svg)](https://pkg.go.dev/iwdgo/filesflood)

# Flooding the file system

Unix main errors on file system are EMFILE, too many files for the process, and ENFILE, too many files for the system.

Module attempts to trigger these errors on Windows and Unix systems.

Repository is related to [Issue 32309](https://github.com/golang/go/issues/32309)

### Usage

Operating system is throttled to ease the demonstration. The module does not revert changes.

**USE WITH CARE**.

CI files for Github/Actions are provided.

`Flooding(timeout int)` is an example of set up where `timeout` in is minutes.

`FilesFloodEmfile()` and `FilesFloodEnfile()` are printing errors sorted by main categories.

### Unix

EMFILE is fairly easy to obtain. ENFILE is never reached.

### Windows

Windows has a virtually limitless number of handle by process (2^24) and 2^16 on 32-bit system.
Handles on a connection is 2^14 as default.
(https://superuser.com/questions/1356320/what-is-the-number-of-open-files-limits)

When using Winsock 2.0, errors must be checked using WSAGetLastError() which is wrapping GetLastError().
(https://stackoverflow.com/questions/15586224/is-wsagetlasterror-just-an-alias-for-getlasterror)
No error has been returned by Windows.

### Build tags

To recover details of the error, Windows uses `syscall.GetLastError()`. This call does not
exist on Unix's. Build tags are used to keep portability.
