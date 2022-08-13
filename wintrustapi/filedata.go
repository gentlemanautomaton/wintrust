package wintrustapi

import (
	"syscall"
	"unsafe"
)

// FileData holds data for wintrust API calls that verify the authenticode
// signature of a file.
type FileData struct {
	UserInterfaceMode    UIMode
	RevocationChecks     RevocationCheck
	File                 FileInfo
	StateAction          StateAction
	StateData            syscall.Handle
	ProviderFlags        uint32
	UserInterfaceContext UIContext
	SignatureSettings    *SignatureSettings
}

type rawFileData struct {
	Size                 uint32
	PolicyCallbackData   uintptr
	SIPClientData        uintptr
	UserInterfaceMode    UIMode
	RevocationChecks     RevocationCheck
	UnionChoice          DataType // Must be DataTypeFile
	File                 *rawFileInfo
	StateAction          StateAction
	StateData            syscall.Handle
	URLReference         *uint16 // Reserved
	ProviderFlags        uint32
	UserInterfaceContext UIContext
	SignatureSettings    *SignatureSettings
}

const rawFileDataSize = uint32(unsafe.Sizeof(rawFileData{}))
