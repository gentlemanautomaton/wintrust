package wintrustapi

import (
	"github.com/gentlemanautomaton/winguid"
	"golang.org/x/sys/windows"
)

// Action is a GUID that identifies an action that can be performed by a
// Windows Trust Provider.
type Action windows.GUID

// Supported trust provider actions.
var (
	DriverVerify    = Action(winguid.New("F750E6C3-38EE-11d1-85E5-00C04FC295EE")) // DRIVER_ACTION_VERIFY
	GenericVerifyV2 = Action(winguid.New("00AAC56B-CD44-11D0-8CC2-00C04FC295EE")) // WINTRUST_ACTION_GENERIC_VERIFY_V2
)
