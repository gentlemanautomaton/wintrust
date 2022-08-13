package wintrustapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modwtsapi32 = windows.NewLazySystemDLL("Wintrust.dll")

	procWinVerifyTrustEx = modwtsapi32.NewProc("WinVerifyTrustEx")
)

// VerifyFile calls the [WinVerifyTrustEx] function within the Windows Trust
// API to verify the authenticode signature of a file. It returns nil if the
// file was signed by a trusted publisher and the signature is valid.
//
// It is the caller's responsibility to call this function twice when
// performing file verification. The first call should use StateActionVerify
// and the second call should use StateActionClose. This is necessary to
// release system resources allocated by the first call.
//
// [WinVerifyTrustEx]: https://docs.microsoft.com/en-us/windows/win32/api/wintrust/nf-wintrust-winverifytrustex
func VerifyFile(hwnd syscall.Handle, action Action, data *FileData) (err error) {
	filePath, err := syscall.UTF16PtrFromString(data.File.Path)
	if err != nil {
		return err
	}

	var fileKnownSubject *windows.GUID
	var emptyGUID windows.GUID
	if data.File.KnownSubject != emptyGUID {
		fileKnownSubject = &data.File.KnownSubject
	}

	rawData := rawFileData{
		Size:              rawFileDataSize,
		UserInterfaceMode: data.UserInterfaceMode,
		RevocationChecks:  data.RevocationChecks,
		UnionChoice:       DataTypeFile,
		File: &rawFileInfo{
			Size:         rawFileInfoSize,
			Path:         filePath,
			Handle:       data.File.Handle,
			KnownSubject: fileKnownSubject,
		},
		StateAction:          data.StateAction,
		StateData:            data.StateData,
		ProviderFlags:        data.ProviderFlags,
		UserInterfaceContext: data.UserInterfaceContext,
		SignatureSettings:    data.SignatureSettings,
	}

	r0, _, e := syscall.SyscallN(
		procWinVerifyTrustEx.Addr(),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&action)),
		uintptr(unsafe.Pointer(&rawData)))
	if r0 != 0 {
		err = syscall.Errno(e)
	} else {
		data.StateData = rawData.StateData
	}
	return
}
