package wintrustapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// FileInfo holds information about a file to be examined by wintrust API
// calls.
type FileInfo struct {
	Path         string
	Handle       syscall.Handle // Optional handle to file with read access
	KnownSubject windows.GUID   // Optional subject type
}

type rawFileInfo struct {
	Size         uint32
	Path         *uint16
	Handle       syscall.Handle // Optional handle to file with read access
	KnownSubject *windows.GUID  // Optional subject type
}

const rawFileInfoSize = uint32(unsafe.Sizeof(rawFileInfo{}))
