package wintrust

import (
	"fmt"
	"syscall"

	"github.com/gentlemanautomaton/wintrust/wintrustapi"
)

// VerifyFile invokes the Windows Trust Verification API to verify the
// authenticode signature of the file with the given path.
//
// It returns nil if the file was signed by a trusted publisher and the
// signature is valid.
func VerifyFile(path string) error {
	data := wintrustapi.FileData{
		UserInterfaceMode: wintrustapi.UserInterfaceNone,
		RevocationChecks:  wintrustapi.RevocationCheckWholeChain,
		File:              wintrustapi.FileInfo{Path: path},
	}

	// The first call retrieves data
	data.StateAction = wintrustapi.StateActionVerify
	err1 := wintrustapi.VerifyFile(syscall.InvalidHandle, wintrustapi.GenericVerifyV2, &data)

	// The second call frees data allocated in the first call
	data.StateAction = wintrustapi.StateActionClose
	err2 := wintrustapi.VerifyFile(syscall.InvalidHandle, wintrustapi.GenericVerifyV2, &data)

	// Check for errors from both calls
	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return fmt.Errorf("failed to release system resources after file signature verification: %v", err2)
	}

	return nil
}
